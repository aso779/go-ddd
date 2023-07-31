package dataspec

import (
	"reflect"
	"testing"

	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
)

func TestInSpecification_IsEmpty(t *testing.T) {
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
				value: []int{42},
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
			r := &InSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSpecification_Query(t *testing.T) {
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
			want: " IN (?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.Query(tt.args.meta); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSpecification_Values(t *testing.T) {
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
				value: []int{42, 45},
			},
			want: []any{[]int{42, 45}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InSpecification{
				field: tt.fields.field,
				value: tt.fields.value,
			}
			if got := r.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
