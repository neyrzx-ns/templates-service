package entity

import (
	"errors"
	"templates-service/internal/domain/template/valueobjects"
)

type Template[Spec valueobjects.Spec] struct {
	id          valueobjects.TemplateID
	name        valueobjects.Name
	description valueobjects.Description
	text        Spec
	variables   []string
}

func (t Template[Spec]) ID() valueobjects.TemplateID {
	return t.id
}

func (t Template[Spec]) Name() valueobjects.Name {
	return t.name
}

func (t Template[Spec]) Description() valueobjects.Description {
	return t.description
}

func (t Template[Spec]) Text() Spec {
	return t.text
}

func NewTemplate[Spec valueobjects.Spec]() Template[Spec] {
	return Template[Spec]{id: valueobjects.NewTemplateID()}
}

type TemplateBuilder[Spec valueobjects.Spec] struct {
	template Template[Spec]
	errs     []error
}

func NewTemplateBuilder[Spec valueobjects.Spec]() TemplateBuilder[Spec] {
	return TemplateBuilder[Spec]{template: NewTemplate[Spec]()}
}

func (b *TemplateBuilder[Spec]) SetName(name string) *TemplateBuilder[Spec] {
	var err error

	if b.template.name, err = valueobjects.NewName(name); err != nil {
		b.errs = append(b.errs, err)
	}

	return b
}

func (b *TemplateBuilder[Spec]) SetDescription(description string) *TemplateBuilder[Spec] {
	var err error
	if b.template.description, err = valueobjects.NewDescription(description); err != nil {
		b.errs = append(b.errs, err)
	}
	return b
}

func (b *TemplateBuilder[Spec]) SetText(text string) *TemplateBuilder[Spec] {
	voText := Spec(text)
	if err := voText.Validate(); err != nil {
		b.errs = append(b.errs, err)
	}

	b.template.text = voText
	return b
}

func (b *TemplateBuilder[Spec]) Build() (Template[Spec], error) {
	if len(b.errs) > 0 {
		return Template[Spec]{}, errors.Join(b.errs...)
	}

	return b.template, nil
}
