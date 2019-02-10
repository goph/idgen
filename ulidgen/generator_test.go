package ulidgen

import (
	"bytes"
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
)

func TestGenerator_New(t *testing.T) {
	clock := clockwork.NewFakeClockAt(time.Unix(1549414354, 0))
	entropy := bytes.NewBufferString("entropy123456789")
	generator := NewGenerator(clock, entropy)

	id, err := generator.New()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := id, "01D304NK2GCNQ78WKFE1WK2CHK"; got != want {
		t.Errorf("the generated id does not match the expected\nexpected: %s\nactual:   %s", want, got)
	}
}
