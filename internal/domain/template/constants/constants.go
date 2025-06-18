package constants

import "regexp"

const (
	NameMaxLength        = 255
	DescriptionMaxLength = 1024
)

// Unicode Categories: https://www.regular-expressions.info/unicode.html#category
var (
	NameAllowedSymbols        = regexp.MustCompile(`^[()\p{L}\p{N}\p{P} ]+$`)
	DescriptionAllowedSymbols = regexp.MustCompile(`^[()\p{L}\p{N}\p{P}\p{C} ]+$`)
)
