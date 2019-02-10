package ulidgen

import (
	"io"
	"time"

	"github.com/oklog/ulid"
	"github.com/pkg/errors"
)

// Clock tells the current time.
type Clock interface {
	// Now tells the current time.
	Now() time.Time
}

// Generator generates a ULID.
type Generator struct {
	clock   Clock
	entropy io.Reader
}

// NewGenerator returns a new Generator.
func NewGenerator(clock Clock, entropy io.Reader) *Generator {
	return &Generator{
		clock:   clock,
		entropy: entropy,
	}
}

// New generates a new ID.
func (g *Generator) New() (string, error) {
	id, err := ulid.New(ulid.Timestamp(g.clock.Now()), g.entropy)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}
