package entrel

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

type ToMany struct {
	Meta               metadata.Meta
	JoinTable          string
	ViaTable           string
	JoinColumns        []JoinColumn
	InverseJoinColumns []JoinColumn
}

func (r ToMany) Join() []string {
	var result []string

	sb := strings.Builder{}
	sb.WriteString("JOIN ")
	sb.WriteString(r.ViaTable)
	sb.WriteString(" ON ")
	for _, v := range r.JoinColumns {
		sb.WriteString(v.Name)
		sb.WriteString("=")
		sb.WriteString(v.ReferencedName)
	}
	result = append(result, sb.String())

	sb = strings.Builder{}
	sb.WriteString("JOIN ")
	sb.WriteString(r.JoinTable)
	sb.WriteString(" ON ")
	for _, v := range r.InverseJoinColumns {
		sb.WriteString(v.Name)
		sb.WriteString("=")
		sb.WriteString(v.ReferencedName)
	}
	result = append(result, sb.String())

	return result
}

func (r ToMany) Table() string {
	return r.JoinTable
}

func (r ToMany) GetMeta() metadata.Meta {
	return r.Meta
}
