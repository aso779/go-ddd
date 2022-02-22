package entrel

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

type ToOne struct {
	Meta        metadata.Meta
	JoinTable   string
	JoinColumns []JoinColumn
}

func (r ToOne) Join() []string {
	sb := strings.Builder{}
	for _, v := range r.JoinColumns {
		sb.WriteString("JOIN ")
		sb.WriteString(r.JoinTable)
		sb.WriteString(" ON ")
		sb.WriteString(v.Name)
		sb.WriteString("=")
		sb.WriteString(r.JoinTable)
		sb.WriteString(".")
		sb.WriteString(v.ReferencedName)
	}
	return []string{sb.String()}
}

func (r ToOne) Prefix() string {
	return r.JoinTable
}

func (r ToOne) GetMeta() metadata.Meta {
	return r.Meta
}
