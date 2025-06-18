package notification

const (
	MaxTextLength = 1024
)

type TelegramType string

func NewTelegramType(text string) (TelegramType, error) {
	// TODO: validation
	return TelegramType(text), nil
}

func (s TelegramType) Validate() error {
	return nil
}
