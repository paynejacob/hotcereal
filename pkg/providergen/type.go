package providergen

import (
	"errors"
	"go/types"
	"strings"
)

type Field struct {
	Name    string
	IsArray bool
	Type    string
}

type Resource struct {
	Name             string
	Package          string
	KeyField         Field
	SearchableFields []Field
	LookupFields     []Field
}

func ResourceFromStruct(name, pkg string, t *types.Struct) (r *Resource, err error) {
	r = &Resource{
		Name: name,
		Package: pkg,
	}

	for i := 0; i < t.NumFields(); i++ {
		fieldAttrs := ParseTag(t.Tag(i))

		field := Field{
			Name:    t.Field(i).Name(),
			IsArray: strings.HasPrefix(t.Field(i).Type().String(), "[]"),
			Type:    getFieldRenderType(t.Field(i)),
		}

		if fieldAttrs.IsKey {
			if r.KeyField.Name != "" {
				err = errors.New(r.Name + ": has more than 1 key field")
				return
			}

			if field.Type != "string" {
				err = errors.New(r.Name + ": key field must be type [string]")
				return
			}

			r.KeyField = field
		}

		if fieldAttrs.Searchable {
			if field.Type != "string" {
				err = errors.New(field.Name + ": searchable fields must be type [string]")
				return
			}

			r.SearchableFields = append(r.SearchableFields, field)
		}

		if fieldAttrs.Lookup {
			r.LookupFields = append(r.LookupFields, field)
		}
	}

	if r.KeyField.Name == "" {
		err = errors.New(r.Name + " does not define a key field")
	}

	return
}

func getFieldRenderType(f *types.Var) string {
	renderType := strings.TrimPrefix(f.Type().String(), "[]")
	customSplit := strings.LastIndexAny(renderType, ".")

	// this indicates a built-in like string or something similar
	if !strings.Contains(renderType, "/") {
		return f.Type().String()
	}

	// TODO: this assumes the type is in the same package
	// if our type comes from the same package just return the type name otherwise prepend the package name
	return renderType[customSplit + 1:]
}
