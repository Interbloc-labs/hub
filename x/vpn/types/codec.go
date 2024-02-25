// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodetypeslegacy "github.com/sentinel-official/hub/x/node/legacy/v1/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypeslegacy "github.com/sentinel-official/hub/x/plan/legacy/v1/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypeslegacy "github.com/sentinel-official/hub/x/provider/legacy/v1/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypeslegacy "github.com/sentinel-official/hub/x/session/legacy/v1/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypeslegacy "github.com/sentinel-official/hub/x/subscription/legacy/v1/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	providertypes.RegisterInterfaces(registry)
	providertypeslegacy.RegisterInterfaces(registry)
	nodetypes.RegisterInterfaces(registry)
	nodetypeslegacy.RegisterInterfaces(registry)
	plantypes.RegisterInterfaces(registry)
	plantypeslegacy.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	subscriptiontypeslegacy.RegisterInterfaces(registry)
	sessiontypes.RegisterInterfaces(registry)
	sessiontypeslegacy.RegisterInterfaces(registry)
}
