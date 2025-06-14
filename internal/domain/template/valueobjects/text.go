package valueobjects

import "errors"

type Spec interface {
	~string

	Validate() error
}

type SMSSpec string

func NewSMSSpec(text string) (SMSSpec, error) {
	// TODO: validation
	return SMSSpec(text), nil
}

func (s SMSSpec) Validate() error {
	return nil
}

type PushNotificationSpec string

func NewPushNotificationSpec(text string) (PushNotificationSpec, error) {
	// TODO: validation
	return PushNotificationSpec(text), nil
}

func (s PushNotificationSpec) Validate() error {
	return errors.New("not implemented")
}

type MessengerSpec string

func NewMessengerSpec(text string) (MessengerSpec, error) {
	// TODO: validation
	return MessengerSpec(text), nil
}

func (s MessengerSpec) Validate() error {
	return nil
}

type EmailSpec string

func NewEmailSpec(text string) (EmailSpec, error) {
	// TODO: validation
	return EmailSpec(text), nil
}

func (s EmailSpec) Validate() error {
	return nil
}
