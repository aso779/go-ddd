package metadata

import "sort"

// Meta entity meta information
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
	//PersistencePresenterMapping returns mapping persistence names to presenter names
	PersistencePresenterMapping() map[string]string
	//Relations returns possible relations
	Relations() map[string]Relation
}

// PrimaryKey struct for primary keys
type PrimaryKey map[string]any

func (r PrimaryKey) Sorted() []PrimaryKey {
	var (
		keys []string
		res  = make([]PrimaryKey, len(r))
	)
	for k, _ := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for k, v := range keys {
		res[k] = PrimaryKey{v: r[v]}
	}

	return res
}

// Keys primary key keys
func (r PrimaryKey) Keys() []string {
	keys := make([]string, len(r))
	i := 0
	for k := range r {
		keys[i] = k
		i++
	}

	return keys
}

func (r PrimaryKey) SortedKeys() []string {
	k := r.Keys()
	sort.Strings(k)

	return k
}

// Values primary key values
func (r PrimaryKey) Values() []any {
	values := make([]any, len(r))
	i := 0
	for _, v := range r {
		values[i] = v
		i++
	}

	return values
}

// IsComposite primary key composite check
func (r PrimaryKey) IsComposite() bool {
	return len(r) > 1
}

// IsEmpty primary key empty check
func (r PrimaryKey) IsEmpty() bool {
	return len(r) == 0
}

// Entity general entity interface
type Entity interface {
	//EntityName entity name
	EntityName() string
	//PrimaryKey entity primary key
	PrimaryKey() PrimaryKey
}

// EntityMetaDecorator entity decorator. Describe entity relations meta data
type EntityMetaDecorator interface {
	Entity() Entity
	Relations
}

// Relations general relations interface
type Relations interface {
	Relations() map[string]Relation
}

// Relation describe infrastructure entity relations
type Relation interface {
	GetMeta() Meta
	Join() []Join
	Table() string
}

type MetaParser func(decorator EntityMetaDecorator) Meta

// EntityMetaContainer describe entity meta container
type EntityMetaContainer interface {
	Add(entity EntityMetaDecorator, parser MetaParser)
	Get(entName string) Meta
}

type Join struct {
	JoinString string
	Args       []any
}

func NewJoin(query string, args ...any) Join {
	return Join{
		JoinString: query,
		Args:       args,
	}
}
