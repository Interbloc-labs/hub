// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/querier
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/client/cli
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/node/client/rest
package node

import (
	"github.com/sentinel-official/hub/x/dvpn/node/client/cli"
	"github.com/sentinel-official/hub/x/dvpn/node/client/rest"
	"github.com/sentinel-official/hub/x/dvpn/node/keeper"
	"github.com/sentinel-official/hub/x/dvpn/node/querier"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

const (
	Codespace              = types.Codespace
	EventTypeSetNode       = types.EventTypeSetNode
	EventTypeUpdateNode    = types.EventTypeUpdateNode
	EventTypeSetNodeStatus = types.EventTypeSetNodeStatus
	AttributeKeyProvider   = types.AttributeKeyProvider
	AttributeKeyAddress    = types.AttributeKeyAddress
	AttributeKeyStatus     = types.AttributeKeyStatus
	AttributeKeyStatusAt   = types.AttributeKeyStatusAt
	ModuleName             = types.ModuleName
	QuerierRoute           = types.QuerierRoute
	CategoryUnknown        = types.CategoryUnknown
	QueryNode              = types.QueryNode
	QueryNodes             = types.QueryNodes
	QueryNodesOfProvider   = types.QueryNodesOfProvider
)

var (
	// functions aliases
	NewKeeper                     = keeper.NewKeeper
	Querier                       = querier.Querier
	RegisterCodec                 = types.RegisterCodec
	ErrorMarshal                  = types.ErrorMarshal
	ErrorUnmarshal                = types.ErrorUnmarshal
	ErrorUnknownMsgType           = types.ErrorUnknownMsgType
	ErrorUnknownQueryType         = types.ErrorUnknownQueryType
	ErrorInvalidField             = types.ErrorInvalidField
	ErrorNoProviderFound          = types.ErrorNoProviderFound
	ErrorDuplicateNode            = types.ErrorDuplicateNode
	ErrorNoNodeFound              = types.ErrorNoNodeFound
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	NodeKey                       = types.NodeKey
	NodeAddressForProviderKey     = types.NodeAddressForProviderKey
	NewMsgRegisterNode            = types.NewMsgRegisterNode
	NewMsgUpdateNode              = types.NewMsgUpdateNode
	NewMsgSetNodeStatus           = types.NewMsgSetNodeStatus
	NodeCategoryFromString        = types.NodeCategoryFromString
	NewQueryNodeParams            = types.NewQueryNodeParams
	NewQueryNodesOfProviderParams = types.NewQueryNodesOfProviderParams
	GetQueryCommands              = cli.GetQueryCommands
	GetTxCommands                 = cli.GetTxCommands
	RegisterRoutes                = rest.RegisterRoutes

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	RouterKey     = types.RouterKey
	StoreKey      = types.StoreKey
	NodeKeyPrefix = types.NodeKeyPrefix
)

type (
	Keeper                     = keeper.Keeper
	GenesisState               = types.GenesisState
	MsgRegisterNode            = types.MsgRegisterNode
	MsgUpdateNode              = types.MsgUpdateNode
	MsgSetNodeStatus           = types.MsgSetNodeStatus
	NodeCategory               = types.NodeCategory
	Node                       = types.Node
	Nodes                      = types.Nodes
	QueryNodeParams            = types.QueryNodeParams
	QueryNodesOfProviderParams = types.QueryNodesOfProviderParams
)
