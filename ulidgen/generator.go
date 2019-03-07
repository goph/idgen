// Package ulidgen generates ULIDs.
package ulidgen

import (
	"io"
	"time"

	"github.com/oklog/ulid"
	"github.com/pkg/errors"
)

// Clock provides the time part of the ULID.
type Clock interface {
	// Now provides the time part of the ULID.
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

// Generate generates a new ID.
func (g *Generator) Generate() (string, error) {
	id, err := ulid.New(ulid.Timestamp(g.clock.Now()), g.entropy)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}
