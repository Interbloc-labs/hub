// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodelegacytypes "github.com/sentinel-official/hub/x/node/legacy/v1/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypeslegacy "github.com/sentinel-official/hub/x/subscription/legacy/v1/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	providertypes.RegisterInterfaces(registry)
	nodetypes.RegisterInterfaces(registry)
	nodelegacytypes.RegisterInterfaces(registry)
	plantypes.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	subscriptiontypeslegacy.RegisterInterfaces(registry)
	sessiontypes.RegisterInterfaces(registry)
}
