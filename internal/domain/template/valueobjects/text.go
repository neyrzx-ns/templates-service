package valueobjects

type Specification interface {
	~string
	Validate() error
}

type SMSSpecification string

func NewSMSSpecification(text string) (SMSSpecification, error) {
	// TODO: validation
	return SMSSpecification(text), nil
}

func (s SMSSpecification) Validate() error {
	return nil
}

type PushNotificationSpecification string

func NewPushNotificationSpecification(text string) (PushNotificationSpecification, error) {
	// TODO: validation
	return PushNotificationSpecification(text), nil
}

func (s PushNotificationSpecification) Validate() error {
	return nil
}

type MessengerSpecification string

func NewMessengerSpecification(text string) (MessengerSpecification, error) {
	// TODO: validation
	return MessengerSpecification(text), nil
}

func (s MessengerSpecification) Validate() error {
	return nil
}

type EmailSpecification string

func NewEmailSpecification(text string) (EmailSpecification, error) {
	// TODO: validation
	return EmailSpecification(text), nil
}

func (s EmailSpecification) Validate() error {
	return nil
}
