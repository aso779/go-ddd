package metadata

//Meta entity meta information
type Meta interface {
	//EntityName returns entity name
	EntityName() string
	//PersistenceName returns persistence name
	PersistenceName() string
	//FieldToPresenter returns presenting name by given field name
	FieldToPresenter(fieldName string) string
	//PresenterToPersistence returns persistence name by given presenter name
	PresenterToPersistence(presenterName string) string
	//PresenterSetToPersistenceSet returns presenter fields name set by given persistence name set
	PresenterSetToPersistenceSet(presenterNames []string) []string
	//PresenterPersistenceMapping returns mapping presenter names to persistence names
	PresenterPersistenceMapping() map[string]string
	//Relations returns possible relations
	Relations() map[string]Relation
}

//Entity general entity interface
type Entity interface {
	//Name entity name
	Name() string
	//PrimaryKey entity primary key
	PrimaryKey() map[string]interface{}
}

//EntityMetaDecorator entity decorator. Describe entity relations meta data
type EntityMetaDecorator interface {
	Entity() Entity
	Relations
}

//Relations general relations interface
type Relations interface {
	Relations() map[string]Relation
}

//Relation describe infrastructure entity relations
type Relation interface {
	Meta() Meta
	Join() []string
	Table() string
}

//EntityMetaContainer describe entity meta container
type EntityMetaContainer interface {
	Add(entity EntityMetaDecorator)
	Get(entName string) Meta
}
