package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField     = errors.Register(ModuleName, 1010, "invalid field")
	ErrorInvalidFrom      = errors.Register(ModuleName, 1020, "invalid from")
	ErrorInvalidProvider  = errors.Register(ModuleName, 1030, "invalid provider")
	ErrorInvalidPrice     = errors.Register(ModuleName, 1040, "invalid price")
	ErrorInvalidRemoteURL = errors.Register(ModuleName, 1050, "invalid remote_url")
	ErrorInvalidStatus    = errors.Register(ModuleName, 1060, "invalid status")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 2010, "provider does not exist")
	ErrorDuplicateNode        = errors.Register(ModuleName, 2020, "duplicate node")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 2030, "node does not exist")
	ErrorInvalidPlanCount     = errors.Register(ModuleName, 2040, "invalid plan count")
)
