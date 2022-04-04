package dataspec

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

const entNameIdx = 1

type Field struct {
	key       string
	entName   string
	fieldName string
}

func NewField(field string) Field {
	fieldParts := strings.Split(field, ".")
	if len(fieldParts) == 1 {
		return Field{
			key:       field,
			entName:   "",
			fieldName: fieldParts[0],
		}
	}

	return Field{
		key:       field,
		entName:   fieldParts[len(fieldParts)-2],
		fieldName: fieldParts[len(fieldParts)-1],
	}
}

func (r *Field) EntName() string {
	return r.entName
}

func (r *Field) FieldName() string {
	return r.fieldName
}

func (r *Field) Key() string {
	return r.key
}

func (r *Field) RelKey() string {
	parts := strings.Split(r.key, ".")

	return strings.Join(parts[entNameIdx:len(parts)-1], ".")
}

//TODO replace to constructor for performance reasons

func (r *Field) RelKeys() []string {
	var key string
	var keys []string
	parts := strings.Split(r.key, ".")

	if len(parts) < 2 {
		return keys
	}

	for i, v := range parts[entNameIdx : len(parts)-1] {
		if i == 0 {
			key = v
			keys = append(keys, key)
		} else {
			key = fmt.Sprintf("%s.%s", key, v)
			keys = append(keys, key)
		}
	}

	return keys
}

//TODO replace to constructor for performance reasons

func (r *Field) ColumnName(meta metadata.Meta) string {
	if r.entName == meta.EntityName() || r.entName == "" {
		table := meta.PersistenceName()
		return fmt.Sprintf("%s.%s", table, meta.PresenterToPersistence(r.fieldName))
	}
	if rel, ok := meta.Relations()[r.RelKey()]; ok {
		return fmt.Sprintf(`"%s"."%s"`, rel.Table(), rel.GetMeta().PresenterToPersistence(r.fieldName))
	}

	//TODO error
	return ""
}
