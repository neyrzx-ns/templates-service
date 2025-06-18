package notification

import "errors"

type InAppType string

func NewInAppType(text string) (InAppType, error) {
	// TODO: validation
	return InAppType(text), nil
}

func (s InAppType) Validate() error {
	if len(s) > 255 {
		return errors.New("push notification specification too long")
	}

	return nil
}
