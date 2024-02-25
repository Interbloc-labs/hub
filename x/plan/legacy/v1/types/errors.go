package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField    = errors.Register(ModuleName, 1010, "invalid field")
	ErrorInvalidFrom     = errors.Register(ModuleName, 1020, "invalid from")
	ErrorInvalidPrice    = errors.Register(ModuleName, 1030, "invalid price")
	ErrorInvalidValidity = errors.Register(ModuleName, 1040, "invalid validity")
	ErrorInvalidBytes    = errors.Register(ModuleName, 1050, "invalid bytes")
	ErrorInvalidId       = errors.Register(ModuleName, 1060, "invalid id")
	ErrorInvalidStatus   = errors.Register(ModuleName, 1070, "invalid status")
	ErrorInvalidAddress  = errors.Register(ModuleName, 1080, "invalid address")
)

var (
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 2010, "provider does not exist")
	ErrorPlanDoesNotExist     = errors.Register(ModuleName, 2020, "plan does not exist")
	ErrorNodeDoesNotExist     = errors.Register(ModuleName, 2030, "node does not exist")
	ErrorUnauthorized         = errors.Register(ModuleName, 2040, "unauthorized")
	DuplicateNodeForPlan      = errors.Register(ModuleName, 2050, "duplicate node for plan")
)
