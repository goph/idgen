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

	testSafeGenerator(t, generator)
}

func TestConstantGenerator_Empty(t *testing.T) {
	generator := &ConstantGenerator{}

	_, err := generator.Generate()
	if err == nil || err.Error() != "no id is configured" {
		t.Errorf("expected an error, got: %s", err)
	}
}

func TestGenerator(t *testing.T) {
	generator := NewGenerator(NewConstantGenerator("id"))

	testGenerator(t, generator)
}

func TestGenerator_Panic(t *testing.T) {
	generator := NewGenerator(NewConstantGenerator(""))

	testGeneratorPanic(t, generator)
}

func TestSafeGeneratorFunc_Generate(t *testing.T) {
	generator := SafeGeneratorFunc(func() (string, error) {
		return "id", nil
	})

	testSafeGenerator(t, generator)
}

func TestGeneratorFunc_Generate(t *testing.T) {
	generator := GeneratorFunc(func() string {
		return "id"
	})

	testGenerator(t, generator)
}

func TestListGenerator(t *testing.T) {
	generator := NewListGenerator([]string{"id"})

	testSafeGenerator(t, generator)

	_, err := generator.Generate()
	if err == nil || err.Error() != "no more ids left" {
		t.Errorf("expected an error, got: %s", err)
	}
}

func TestListGenerator_Empty(t *testing.T) {
	generator := &ListGenerator{}

	_, err := generator.Generate()
	if err == nil || err.Error() != "no more ids left" {
		t.Errorf("expected an error, got: %s", err)
	}
}

func testSafeGenerator(t *testing.T, generator SafeGenerator) {
	id, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if id != "id" {
		t.Errorf("expected the value \"id\", got: %s", id)
	}
}

func testGenerator(t *testing.T, generator Generator) {
	id := generator.Generate()

	if id != "id" {
		t.Errorf("expected the value \"id\", got: %s", id)
	}
}

func testGeneratorPanic(t *testing.T, generator Generator) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("Generator is expected to fail")
		}
	}()

	_ = generator.Generate()
}
