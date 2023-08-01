package dataspec

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

func join(meta metadata.Meta, field Field) []metadata.Join {
	var joins []metadata.Join

	if rel, ok := meta.Relations()[field.EntName()]; ok {
		joins = append(joins, rel.Join()...)
	}

	for _, key := range field.RelKeys() {
		if rel, ok := meta.Relations()[key]; ok {
			joins = append(joins, rel.Join()...)
		}
	}

	return joins
}
