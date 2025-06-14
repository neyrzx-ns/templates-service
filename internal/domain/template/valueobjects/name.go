package valueobjects

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	// TODO: validation
	return Name{value: value}, nil
}
