package entity

import (
	"regexp"
	"templates-service/internal/domain/template/valueobjects"
	"templates-service/internal/domain/template/valueobjects/notification"
)

var (
	templateVariablesGroupName = "variable"
	templateVariablesRegexp    = regexp.MustCompile(`{{\s?.(?P<variable>\w+)\s?}}`)
)

type Template[NotificationType notification.Type] struct {
	id           valueobjects.TemplateID
	name         valueobjects.Name
	description  valueobjects.Description
	notification NotificationType
	variables    []string
}

func (t Template[NotificationType]) ID() valueobjects.TemplateID {
	return t.id
}

func (t Template[NotificationType]) Name() valueobjects.Name {
	return t.name
}

func (t Template[NotificationType]) Description() valueobjects.Description {
	return t.description
}

func (t Template[NotificationType]) Text() NotificationType {
	return t.notification
}

func NewTemplate[T notification.Type]() Template[T] {
	return Template[T]{id: valueobjects.NewTemplateID()}
}
