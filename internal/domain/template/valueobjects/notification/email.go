package notification

type EmailType string

func NewEmailType(text string) (EmailType, error) {
	// TODO: validation
	return EmailType(text), nil
}

func (s EmailType) Validate() error {
	return nil
}
