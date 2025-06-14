package valueobjects

import "github.com/google/uuid"

type TemplateGroupID uuid.UUID

func NewTemplateGroupID() TemplateGroupID {
	uid, _ := uuid.NewUUID()
	return TemplateGroupID(uid)
}

func (id TemplateGroupID) String() string {
	return uuid.UUID(id).String()
}
