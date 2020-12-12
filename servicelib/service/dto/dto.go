package dto

import (
	"reflect"

	"github.com/Meduzz/func-lib/servicelib/service/annotation"
)

type (
	EntityDTO struct {
		Name        string      `json:"name"`
		Fields      []*FieldDTO `json:"fields"`
		Description string      `json:"description,omitempty"`
	}

	FieldDTO struct {
		Name        string                  `json:"name"`
		Fields      []*FieldDTO             `json:"fields,omitempty"`
		Annotations []annotation.Annotation `json:"annotations,omitempty"`
		Description string                  `json:"description,omitempty"`
	}
)

func NewEntity(name string, fields []*FieldDTO) *EntityDTO {
	return &EntityDTO{
		Name:   name,
		Fields: fields,
	}
}

func FromStruct(dto interface{}) *EntityDTO {
	entity := reflect.TypeOf(dto)

	if entity.Kind() == reflect.Ptr {
		entity = entity.Elem()
	}

	fields := make([]*FieldDTO, 0)

	for i := 0; i < entity.NumField(); i++ {
		field := entity.Field(i)
		if name, ok := field.Tag.Lookup("json"); ok {
			if name != "" {
				f := NewField(name, annotation.Type(field.Type.String()))
				fields = append(fields, f)
			}
		}
	}

	return NewEntity(entity.Name(), fields)
}

func NewField(name string, annotations ...annotation.Annotation) *FieldDTO {
	return &FieldDTO{
		Name:        name,
		Annotations: annotations,
	}
}

func NewFieldWithChildren(name string, fields []*FieldDTO) *FieldDTO {
	return &FieldDTO{
		Name:   name,
		Fields: fields,
	}
}

func Fields(fields ...*FieldDTO) []*FieldDTO {
	return fields
}

func (f *FieldDTO) SetDescription(text string) {
	f.Description = text
}

func (f *FieldDTO) AddAnnotation(annotation annotation.Annotation) {
	f.Annotations = append(f.Annotations, annotation)
}

func (e *EntityDTO) SetDescription(text string) {
	e.Description = text
}
