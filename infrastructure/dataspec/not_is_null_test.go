package dataspec

import (
	"reflect"
	"testing"

	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
)

func TestNotIsNullSpecification_IsEmpty(t *testing.T) {
	type fields struct {
		field Field
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
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NotIsNullSpecification{
				field: tt.fields.field,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIsNullSpecification_Query(t *testing.T) {
	type fields struct {
		field Field
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
			},
			args: args{
				meta: entmeta.NewMeta(),
			},
			want: " IS NOT NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NotIsNullSpecification{
				field: tt.fields.field,
			}
			if got := r.Query(tt.args.meta); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotIsNullSpecification_Values(t *testing.T) {
	type fields struct {
		field Field
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
			},
			want: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NotIsNullSpecification{
				field: tt.fields.field,
			}
			if got := r.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
