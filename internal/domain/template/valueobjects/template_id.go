package valueobjects

import "github.com/google/uuid"

type TemplateID uuid.UUID

func NewTemplateID() TemplateID {
	uid, _ := uuid.NewUUID()
	return TemplateID(uid)
}

func (id TemplateID) String() string {
	return uuid.UUID(id).String()
}
