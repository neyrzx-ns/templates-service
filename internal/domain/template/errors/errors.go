package errors

import (
	"errors"
	"fmt"

	"templates-service/internal/domain/template/constants"
)

var (
	ErrNameEmpty                 = errors.New("name couldn't be empty")
	ErrNameLength                = errors.New(fmt.Sprintf("name could be less or equal to %d", constants.NameMaxLength))
	ErrNameAllowedSymbols        = errors.New("name contains restricted symbols")
	ErrDescriptionEmpty          = errors.New("description couldn't be empty")
	ErrDescriptionLength         = errors.New(fmt.Sprintf("description could be less or equal to %d", constants.DescriptionMaxLength))
	ErrDescriptionAllowedSymbols = errors.New("description contains restricted symbols")
)
