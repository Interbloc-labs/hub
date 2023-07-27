package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

// The following line asserts that the `msgServer` type implements the `types.MsgServiceServer` interface.
var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `types.MsgServiceServer` interface.
type msgServer struct {
	Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `types.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

// MsgCancel cancels an active subscription.
// It validates the cancel request and performs necessary operations to set the subscription to the inactive state.
func (k *msgServer) MsgCancel(c context.Context, msg *types.MsgCancelRequest) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the subscription from the Store based on the provided subscription ID (msg.ID).
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription is in an active state. If not, return an error.
	if !subscription.GetStatus().Equal(hubtypes.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.GetID(), subscription.GetStatus())
	}

	// Check if the `msg.From` address matches the owner address of the subscription. If not, return an error.
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Delete the subscription from the Store for the time it becomes inactive.
	k.DeleteSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// Calculate the duration for which the subscription will be in the inactive state.
	statusChangeDelay := k.StatusChangeDelay(ctx)
	subscription.SetInactiveAt(
		subscription.GetInactiveAt().Add(statusChangeDelay),
	)
	subscription.SetStatus(hubtypes.StatusInactivePending)
	subscription.SetStatusAt(ctx.BlockTime())

	// Update the subscription in the Store.
	k.SetSubscription(ctx, subscription)

	// Add the subscription back to the Store with the new inactive time.
	k.SetSubscriptionForInactiveAt(ctx, subscription.GetInactiveAt(), subscription.GetID())

	// Emit an event to notify that the subscription status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdateStatus{
			ID:     subscription.GetID(),
			Status: subscription.GetStatus(),
		},
	)

	// Check if there is an associated payout for this subscription and remove it.
	payout, found := k.GetPayout(ctx, subscription.GetID())
	if found {
		k.DeletePayoutForNextAt(ctx, payout.NextAt, payout.ID)

		// Reset the `NextAt` field of the payout and update it in the Store.
		payout.NextAt = time.Time{}
		k.SetPayout(ctx, payout)
	}

	return &types.MsgCancelResponse{}, nil
}

// MsgAllocate allocates bandwidth to another address.
// It validates the allocation request and updates the granted bytes for both the sender and receiver of the bandwidth.
func (k *msgServer) MsgAllocate(c context.Context, msg *types.MsgAllocateRequest) (*types.MsgAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the subscription from the Store based on the provided subscription ID (msg.ID).
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}

	// Check if the subscription type is a plan. If not, return an error.
	if subscription.Type() != types.TypePlan {
		return nil, types.NewErrorInvalidSubscriptionType(subscription.GetID(), subscription.Type())
	}

	// Check if the `msg.From` address matches the owner address of the subscription. If not, return an error.
	if !fromAddr.Equals(subscription.GetAddress()) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	// Get the existing allocation for the sender.
	fromAlloc, found := k.GetAllocation(ctx, subscription.GetID(), fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.GetID(), fromAddr)
	}

	// Convert the `msg.Address` (receiver's address) from Bech32 format to an `sdk.AccAddress`.
	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Get the existing allocation for the receiver.
	toAlloc, found := k.GetAllocation(ctx, subscription.GetID(), toAddr)
	if !found {
		// If the receiver has no existing allocation, create a new one.
		toAlloc = types.Allocation{
			ID:            subscription.GetID(),
			Address:       toAddr.String(),
			GrantedBytes:  sdk.ZeroInt(),
			UtilisedBytes: sdk.ZeroInt(),
		}

		// Update the subscription in the Store to associate it with the new receiver.
		k.SetSubscriptionForAccount(ctx, toAddr, subscription.GetID())
	}

	// Calculate the available bytes for the sender and check if it is sufficient for the allocation.
	grantedBytes := fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
	utilisedBytes := fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
	availableBytes := grantedBytes.Sub(utilisedBytes)

	if msg.Bytes.GT(availableBytes) {
		return nil, types.NewErrorInsufficientBytes(subscription.GetID(), msg.Bytes)
	}

	// Update the allocation for the sender after deducting the allocated bytes.
	fromAlloc.GrantedBytes = availableBytes.Sub(msg.Bytes)
	if fromAlloc.GrantedBytes.LT(fromAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), fromAddr)
	}

	// Update the sender's allocation in the Store.
	k.SetAllocation(ctx, fromAlloc)

	// Emit an event to notify that the sender's allocation has been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: fromAlloc.Address,
			Bytes:   fromAlloc.GrantedBytes,
		},
	)

	// Update the allocation for the receiver after adding the allocated bytes.
	toAlloc.GrantedBytes = msg.Bytes
	if toAlloc.GrantedBytes.LT(toAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.GetID(), toAddr)
	}

	// Update the receiver's allocation in the Store.
	k.SetAllocation(ctx, toAlloc)

	// Emit an event to notify that the receiver's allocation has been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventAllocate{
			ID:      subscription.GetID(),
			Address: toAlloc.Address,
			Bytes:   toAlloc.GrantedBytes,
		},
	)

	return &types.MsgAllocateResponse{}, nil
}
