package entmeta

import "github.com/aso779/go-ddd/domain/usecase/metadata"

type Container struct {
	cache map[string]metadata.Meta
}

func NewContainer() *Container {
	return &Container{
		cache: make(map[string]metadata.Meta),
	}
}

func (r *Container) Add(decorator metadata.EntityMetaDecorator, parser metadata.MetaParser) {
	r.cache[decorator.Entity().Name()] = parser(decorator)
}

func (r *Container) Get(entName string) metadata.Meta {
	return r.cache[entName]
}
