package idgen

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	id := Must("id", nil)
	if id != "id" {
		t.Errorf("expected the value \"id\", got: %s", id)
	}
}

func TestMust_Panic(t *testing.T) {
	var err = errors.New("failed to generate ID")

	defer func() {
		e := recover()
		if e == nil {
			t.Fatal("Must is expected to fail")
		}

		if e != err {
			t.Errorf("Must is expected to fail with error\nexpected: %s\nactual:   %s", err, e)
		}
	}()

	Must("", err)
}

func TestConstantGenerator(t *testing.T) {
	generator := NewConstantGenerator("id")

	id, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if id != "id" {
		t.Errorf("expected the value \"id\", got: %s", id)
	}
}

func TestConstantGenerator_Empty(t *testing.T) {
	generator := &ConstantGenerator{}

	_, err := generator.Generate()
	if err == nil || err.Error() != "no id is configured" {
		t.Errorf("expected an error, got: %s", err)
	}
}

func TestMustGenerator_Panic(t *testing.T) {
	generator := NewMustGenerator(NewConstantGenerator("id"))

	id := generator.Generate()

	if id != "id" {
		t.Errorf("expected the value \"id\", got: %s", id)
	}
}

func TestMustGenerator(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("MustGenerator.Generate is expected to fail")
		}
	}()

	generator := NewMustGenerator(NewConstantGenerator(""))

	_ = generator.Generate()
}
