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

type metaParser func(decorator metadata.EntityMetaDecorator) metadata.Meta

func (r *Container) AddEntityDecorator(decorator metadata.EntityMetaDecorator, parser metaParser) {
	r.cache[decorator.Entity().Name()] = parser(decorator)
}

func (r *Container) GetCache(entName string) metadata.Meta {
	return r.cache[entName]
}
