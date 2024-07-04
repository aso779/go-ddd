package dataspec

import (
	"reflect"
	"testing"

	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
)

func TestEqualSpecification_IsEmpty(t *testing.T) {
	type fields struct {
		field Field
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is not empty",
			fields: fields{
				field: Field{
					key:       "TestKey",
					entName:   "TestEntity",
					fieldName: "TestField",
				},
				value: "TestValue",
			},
			want: false,
		},
		{
			name: "is empty",
			fields: fields{
				field: Field{
					key:       "TestKey",
					entName:   "TestEntity",
					fieldName: "TestField",
				},
				value: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EqualSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualSpecification_Joins(t *testing.T) {
	type fields struct {
		field Field
		value any
	}
	type args struct {
		meta metadata.Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []metadata.Join
	}{
		{
			name: "join",
			fields: fields{
				field: Field{
					key:       "RelatedEntity.relId",
					entName:   "RelatedEntity",
					fieldName: "relId",
				},
				value: "TestValue",
			},
			args: args{
				meta: func() metadata.Meta {
					meta := entmeta.NewMeta()
					meta.SetEntityName("RootEntity")
					meta.SetPersistenceName("root_entity")
					meta.AddPresenterToPersistence("id", "id")
					meta.AddFieldToPresenter("ID", "id")
					meta.AddPersistenceToPresenter("id", "id")
					meta.AddPresenterToPersistence("relId", "rel_id")
					meta.AddFieldToPresenter("RelID", "relId")
					meta.AddPersistenceToPresenter("rel_id", "relId")

					relationMeta := entmeta.NewMeta()
					relationMeta.SetEntityName("RelatedEntity")
					relationMeta.SetPersistenceName("related_entity")

					relations := map[string]metadata.Relation{}

					meta.SetRelations(relations)

					return meta
				}(),
			},
			want: []metadata.Join{
				{
					JoinString: "JOIN related_table ON rel_id=related_table.id",
					Args:       nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EqualSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.Joins(tt.args.meta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Joins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualSpecification_Query(t *testing.T) {
	type fields struct {
		field Field
		value any
	}
	type args struct {
		meta metadata.Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "operand",
			fields: fields{
				field: Field{
					key:       "RootEntity.id",
					entName:   "RootEntity",
					fieldName: "id",
				},
				value: "TestValue",
			},
			args: args{
				meta: entmeta.NewMeta(),
			},
			want: " = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EqualSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.Query(tt.args.meta); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualSpecification_Values(t *testing.T) {
	type fields struct {
		field Field
		value any
	}
	tests := []struct {
		name   string
		fields fields
		want   []any
	}{
		{
			name: "value",
			fields: fields{
				field: Field{
					key:       "RootEntity.id",
					entName:   "RootEntity",
					fieldName: "id",
				},
				value: "TestValue",
			},
			want: []any{"TestValue"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &EqualSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
