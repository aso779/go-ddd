package metadata

import "testing"

var pk = PrimaryKey{"id": 123}
var compositePk = PrimaryKey{"someId": 123, "otherId": 321, "stringId": "qwerty"}
var emptyPk = PrimaryKey{}

func TestPrimaryKey_Keys(t *testing.T) {
	if len(pk.Keys()) != 1 {
		t.Error("Mismatch keys length")
	}
}

func TestPrimaryKey_CompositeKeys(t *testing.T) {
	if len(compositePk.Keys()) != 3 {
		t.Error("Mismatch keys length")
	}
}

func TestPrimaryKey_Values(t *testing.T) {
	if len(pk.Values()) != 1 {
		t.Error("Mismatch keys length")
	}
}

func TestPrimaryKey_CompositeValues(t *testing.T) {
	if len(compositePk.Values()) != 3 {
		t.Error("Mismatch keys length")
	}
}

func TestPrimaryKey_IsComposite(t *testing.T) {
	if pk.IsComposite() {
		t.Error("Mismatch pk type")
	}
}

func TestPrimaryKey_CompositeIsComposite(t *testing.T) {
	if !compositePk.IsComposite() {
		t.Error("Mismatch pk type")
	}
}

func TestPrimaryKey_IsEmpty(t *testing.T) {
	if !emptyPk.IsEmpty() {
		t.Error("PK not empty")
	}
}

func TestPrimaryKey_NotIsEmpty(t *testing.T) {
	if pk.IsEmpty() {
		t.Error("PK empty")
	}
}
