package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField          = errors.Register(ModuleName, 1010, "invalid field")
	ErrorInvalidFrom           = errors.Register(ModuleName, 1020, "invalid from")
	ErrorInvalidId             = errors.Register(ModuleName, 1030, "invalid id")
	ErrorInvalidNode           = errors.Register(ModuleName, 1040, "invalid node")
	ErrorInvalidProofId        = errors.Register(ModuleName, 1050, "invalid proof->id")
	ErrorInvalidProofDuration  = errors.Register(ModuleName, 1060, "invalid proof->duration")
	ErrorInvalidProofBandwidth = errors.Register(ModuleName, 1070, "invalid proof->bandwidth")
	ErrorInvalidSignature      = errors.Register(ModuleName, 1080, "invalid signature")
	ErrorInvalidRating         = errors.Register(ModuleName, 1090, "invalid rating")
)

var (
	ErrorSubscriptionDoesNotExit   = errors.Register(ModuleName, 2010, "subscription does not exist")
	ErrorInvalidSubscriptionStatus = errors.Register(ModuleName, 2020, "invalid subscription status")
	ErrorUnauthorized              = errors.Register(ModuleName, 2030, "unauthorized")
	ErrorQuotaDoesNotExist         = errors.Register(ModuleName, 2040, "quota does not exist")
	ErrorFailedToVerifyProof       = errors.Register(ModuleName, 2050, "failed to verify proof")
	ErrorNotEnoughQuota            = errors.Register(ModuleName, 2060, "not enough quota")
	ErrorNodeDoesNotExist          = errors.Register(ModuleName, 2070, "node does not exist")
	ErrorInvalidNodeStatus         = errors.Register(ModuleName, 2080, "invalid node status")
	ErrorSessionDoesNotExist       = errors.Register(ModuleName, 2090, "session does not exist")
	ErrorInvalidSessionStatus      = errors.Register(ModuleName, 2100, "invalid session status")
	ErrorNodeAddressMismatch       = errors.Register(ModuleName, 2110, "node address mismatch")
	ErrorNodeDoesNotExistForPlan   = errors.Register(ModuleName, 2120, "node does not exist for plan")
	ErrorDuplicateSession          = errors.Register(ModuleName, 2130, "duplicate session")
)
