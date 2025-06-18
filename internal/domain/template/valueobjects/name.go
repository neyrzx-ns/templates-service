package valueobjects

import (
	"templates-service/internal/domain/template/constants"
	"templates-service/internal/domain/template/errors"
)

type Name struct {
	value string
}

var NameNil = Name{}

func NewName(value string) (Name, error) {
	if len(value) == 0 {
		return NameNil, errors.ErrNameEmpty
	}

	if len(value) > constants.NameMaxLength {
		return NameNil, errors.ErrNameLength
	}

	if !constants.NameAllowedSymbols.Match([]byte(value)) {
		return NameNil, errors.ErrNameAllowedSymbols
	}

	return Name{value: value}, nil
}
