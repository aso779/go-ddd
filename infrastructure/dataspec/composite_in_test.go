package dataspec

import (
	"reflect"
	"testing"

	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

func TestCompositeInSpecification_IsEmpty(t *testing.T) {
	type fields struct {
		fields []Field
		values any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompositeInSpecification{
				fields: tt.fields.fields,
				values: tt.fields.values,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompositeInSpecification_Joins(t *testing.T) {
	type fields struct {
		fields []Field
		values any
	}
	type args struct {
		in0 metadata.Meta
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompositeInSpecification{
				fields: tt.fields.fields,
				values: tt.fields.values,
			}
			if got := r.Joins(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Joins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompositeInSpecification_Query(t *testing.T) {
	type fields struct {
		fields []Field
		values any
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompositeInSpecification{
				fields: tt.fields.fields,
				values: tt.fields.values,
			}
			if got := r.Query(tt.args.meta); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompositeInSpecification_Values(t *testing.T) {
	type fields struct {
		fields []Field
		values any
	}
	tests := []struct {
		name   string
		fields fields
		want   []any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CompositeInSpecification{
				fields: tt.fields.fields,
				values: tt.fields.values,
			}
			if got := r.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCompositeIn(t *testing.T) {
	type args struct {
		fields []string
		values any
	}
	tests := []struct {
		name string
		args args
		want dataset.Specifier
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompositeIn(tt.args.fields, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompositeIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
