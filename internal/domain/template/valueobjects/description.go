package valueobjects

type Description struct {
	value string
}

func NewDescription(description string) (Description, error) {
	// TODO: validation
	return Description{description}, nil
}
