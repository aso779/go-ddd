package entmeta

import "github.com/aso779/go-ddd/domain/usecase/metadata"

type Meta struct {
	decorator              metadata.EntityMetaDecorator
	entityName             string
	persistenceName        string
	relations              map[string]metadata.Relation
	fieldToPresenter       map[string]string
	presenterToPersistence map[string]string
	persistenceToPresenter map[string]string
}

func NewMeta() *Meta {
	return &Meta{
		relations:              make(map[string]metadata.Relation),
		fieldToPresenter:       make(map[string]string),
		presenterToPersistence: make(map[string]string),
		persistenceToPresenter: make(map[string]string),
	}
}

func (r *Meta) SetDecorator(decorator metadata.EntityMetaDecorator) {
	r.decorator = decorator
}

func (r *Meta) SetEntityName(name string) {
	r.entityName = name
}

func (r *Meta) SetPersistenceName(name string) {
	r.persistenceName = name
}

func (r *Meta) SetRelations(relations map[string]metadata.Relation) {
	r.relations = relations
}

func (r *Meta) AddFieldToPresenter(fieldName string, presenterName string) {
	r.fieldToPresenter[fieldName] = presenterName
}

func (r *Meta) AddPresenterToPersistence(presenterName string, persistenceName string) {
	r.presenterToPersistence[presenterName] = persistenceName
}

func (r *Meta) AddPersistenceToPresenter(dbName string, jsonName string) {
	r.persistenceToPresenter[dbName] = jsonName
}

func (r *Meta) EntityName() string {
	return r.entityName
}

func (r *Meta) PersistenceName() string {
	return r.persistenceName
}

func (r *Meta) FieldToPresenter(fieldName string) string {
	return r.fieldToPresenter[fieldName]
}

func (r *Meta) PresenterToPersistence(presenterName string) string {
	return r.presenterToPersistence[presenterName]
}

func (r *Meta) PresenterSetToPersistenceSet(presenterNames []string) []string {
	var dbFields []string
	for _, v := range presenterNames {
		dbFields = append(dbFields, r.presenterToPersistence[v])
	}
	return dbFields
}

func (r *Meta) PresenterPersistenceMapping() map[string]string {
	return r.presenterToPersistence
}

func (r *Meta) Relations() map[string]metadata.Relation {
	return r.relations
}
