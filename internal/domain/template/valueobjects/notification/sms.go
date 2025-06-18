package notification

type SMSType string

func NewSMSType(text string) (SMSType, error) {
	// TODO: validation
	return SMSType(text), nil
}

func (s SMSType) Validate() error {
	return nil
}
