// Package ulidgen generates ULIDs.
package ulidgen

import (
	"io"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/pkg/errors"
)

// Clock provides the time part of the ULID.
type Clock interface {
	// Now provides the time part of the ULID.
	Now() time.Time
}

type systemClock struct{}

func (c *systemClock) Now() time.Time {
	return time.Now()
}

// Generator generates a ULID.
type Generator struct {
	clock   Clock
	entropy io.Reader
}

// Option configures a Generator using the functional options paradigm popularized by Rob Pike and Dave Cheney.
// If you're unfamiliar with this style,
// see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html and
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis.
type Option interface {
	apply(g *Generator)
}

type optionFunc func(*Generator)

func (f optionFunc) apply(g *Generator) { f(g) }

// TimeSource configures a time source of a generator.
func TimeSource(clock Clock) Option {
	return optionFunc(func(g *Generator) {
		g.clock = clock
	})
}

// EntropySource configures an entropy source of a generator.
func EntropySource(entropy io.Reader) Option {
	return optionFunc(func(g *Generator) {
		g.entropy = entropy
	})
}

// NewGenerator returns a new Generator.
func NewGenerator(opts ...Option) *Generator {
	g := &Generator{}

	for _, opt := range opts {
		opt.apply(g)
	}

	// Default time source
	if g.clock == nil {
		g.clock = &systemClock{}
	}

	// Default entropy source
	if g.entropy == nil {
		g.entropy = ulid.Monotonic(rand.New(rand.NewSource(g.clock.Now().UnixNano())), 0)
	}

	return g
}

// Generate generates a new ID.
func (g *Generator) Generate() (string, error) {
	id, err := ulid.New(ulid.Timestamp(g.clock.Now()), g.entropy)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}
