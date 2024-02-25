package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField   = errors.Register(ModuleName, 1010, "invalid field")
	ErrorInvalidFrom    = errors.Register(ModuleName, 1020, "invalid from")
	ErrorInvalidAddress = errors.Register(ModuleName, 1030, "invalid address")
	ErrorInvalidDeposit = errors.Register(ModuleName, 1040, "invalid deposit")
	ErrorInvalidId      = errors.Register(ModuleName, 1050, "invalid id")
	ErrorInvalidDenom   = errors.Register(ModuleName, 1060, "invalid denom")
	ErrorInvalidBytes   = errors.Register(ModuleName, 1070, "invalid bytes")
)

var (
	ErrorPlanDoesNotExist          = errors.Register(ModuleName, 2010, "plan does not exist")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 2020, "node does not exist")
	ErrorUnauthorized              = errors.Register(ModuleName, 2030, "unauthorized")
	ErrorInvalidPlanStatus         = errors.Register(ModuleName, 2040, "invalid plan status")
	ErrorPriceDoesNotExist         = errors.Register(ModuleName, 2050, "price does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 2060, "invalid node status")
	ErrorSubscriptionDoesNotExist  = errors.Register(ModuleName, 2070, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 2080, "invalid subscription status")
	ErrorCanNotSubscribe           = errors.Register(ModuleName, 2090, "can not subscribe")
	ErrorInvalidQuota              = errors.Register(ModuleName, 2100, "invalid quota")
	ErrorDuplicateQuota            = errors.Register(ModuleName, 2110, "duplicate quota")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 2120, "quota does not exist")
	ErrorCanNotAddQuota            = errors.Register(ModuleName, 2130, "can not add quota")
	ErrorCanNotUpdateQuota         = errors.Register(ModuleName, 2140, "can not update quota")
	ErrorCanNotCancel              = errors.Register(ModuleName, 2150, "can not cancel")
)
