package valueobjects

import (
	"templates-service/internal/domain/template/constants"
	"templates-service/internal/domain/template/errors"
)

type Description struct {
	value string
}

func NewDescription(value string) (Description, error) {
	if len(value) == 0 {
		return Description{}, errors.ErrDescriptionEmpty
	}

	if len(value) > constants.DescriptionMaxLength {
		return Description{}, errors.ErrDescriptionLength
	}

	if !constants.DescriptionAllowedSymbols.Match([]byte(value)) {
		return Description{}, errors.ErrDescriptionAllowedSymbols
	}

	return Description{value}, nil
}
