package entity

import (
	"errors"

	"templates-service/internal/domain/template/valueobjects"
	"templates-service/internal/domain/template/valueobjects/notification"
)

type TemplateBuilder[NotificationType notification.Type] struct {
	template Template[NotificationType]
	errs     []error
}

func NewTemplateBuilder[NotificationType notification.Type]() TemplateBuilder[NotificationType] {
	return TemplateBuilder[NotificationType]{template: NewTemplate[NotificationType]()}
}

func (b *TemplateBuilder[NotificationType]) SetName(name string) *TemplateBuilder[NotificationType] {
	var err error

	if b.template.name, err = valueobjects.NewName(name); err != nil {
		b.errs = append(b.errs, err)
	}

	return b
}

func (b *TemplateBuilder[NotificationType]) SetDescription(description string) *TemplateBuilder[NotificationType] {
	var err error
	if b.template.description, err = valueobjects.NewDescription(description); err != nil {
		b.errs = append(b.errs, err)
	}
	return b
}

func (b *TemplateBuilder[NotificationType]) SetText(text string) *TemplateBuilder[NotificationType] {
	voText := NotificationType(text)
	if err := voText.Validate(); err != nil {
		b.errs = append(b.errs, err)
	}

	groupNameIndex := templateVariablesRegexp.SubexpIndex(templateVariablesGroupName)

	for _, variable := range templateVariablesRegexp.FindAllStringSubmatch(text, -1) {
		b.template.variables = append(b.template.variables, variable[groupNameIndex])
	}

	b.template.notification = voText

	return b
}

func (b *TemplateBuilder[NotificationType]) Build() (Template[NotificationType], error) {
	if len(b.errs) > 0 {
		return Template[NotificationType]{}, errors.Join(b.errs...)
	}

	return b.template, nil
}
