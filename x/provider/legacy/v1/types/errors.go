package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField       = errors.Register(ModuleName, 1010, "invalid field")
	ErrorInvalidFrom        = errors.Register(ModuleName, 1020, "invalid from")
	ErrorInvalidName        = errors.Register(ModuleName, 1030, "invalid name")
	ErrorInvalidIdentity    = errors.Register(ModuleName, 1040, "invalid identity")
	ErrorInvalidWebsite     = errors.Register(ModuleName, 1050, "invalid website")
	ErrorInvalidDescription = errors.Register(ModuleName, 1060, "invalid description")
)

var (
	ErrorDuplicateProvider    = errors.Register(ModuleName, 2010, "duplicate provider")
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 2020, "provider does not exist")
)
