package ulidgen

import (
	"bytes"
	"testing"
	"time"
)

type fakeClock struct {
	now time.Time
}

func (c *fakeClock) Now() time.Time {
	return c.now
}

func TestNewGenerator(t *testing.T) {
	generator := NewGenerator()

	id, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Errorf("the generator is expected to work with default options")
	}
}

func TestGenerator_Generate(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1549414354, 0)}
	entropy := bytes.NewBufferString("entropy123456789")
	generator := NewGenerator(TimeSource(clock), EntropySource(entropy))

	id, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := id, "01D304NK2GCNQ78WKFE1WK2CHK"; got != want {
		t.Errorf("the generated id does not match the expected\nexpected: %s\nactual:   %s", want, got)
	}
}
