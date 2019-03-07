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

func TestGenerator_New(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1549414354, 0)}
	entropy := bytes.NewBufferString("entropy123456789")
	generator := NewGenerator(clock, entropy)

	id, err := generator.Generate()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := id, "01D304NK2GCNQ78WKFE1WK2CHK"; got != want {
		t.Errorf("the generated id does not match the expected\nexpected: %s\nactual:   %s", want, got)
	}
}
