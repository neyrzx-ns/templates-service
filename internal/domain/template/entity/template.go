package entity

import (
	"errors"
	"regexp"
	"templates-service/internal/domain/template/valueobjects"
)

var (
	templateVariablesGroupName = "variable"
	templateVariablesRegexp    = regexp.MustCompile(`{{\s?.(?P<variable>\w+)\s?}}`)
)

type Template[Specification valueobjects.Specification] struct {
	id          valueobjects.TemplateID
	name        valueobjects.Name
	description valueobjects.Description
	text        Specification
	variables   []string
}

func (t Template[Specification]) ID() valueobjects.TemplateID {
	return t.id
}

func (t Template[Specification]) Name() valueobjects.Name {
	return t.name
}

func (t Template[Specification]) Description() valueobjects.Description {
	return t.description
}

func (t Template[Specification]) Text() Specification {
	return t.text
}

func NewTemplate[Spec valueobjects.Specification]() Template[Spec] {
	return Template[Spec]{id: valueobjects.NewTemplateID()}
}

type TemplateBuilder[Spec valueobjects.Specification] struct {
	template Template[Spec]
	errs     []error
}

func NewTemplateBuilder[Spec valueobjects.Specification]() TemplateBuilder[Spec] {
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

	groupNameIndex := templateVariablesRegexp.SubexpIndex(templateVariablesGroupName)

	for _, variable := range templateVariablesRegexp.FindAllStringSubmatch(text, -1) {
		b.template.variables = append(b.template.variables, variable[groupNameIndex])
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
