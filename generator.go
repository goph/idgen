package idgen

import (
	"errors"
)

// Generator generates a new ID.
type Generator interface {
	// New generates a new ID.
	New() (string, error)
}

// Must panics if a generator function returns with an error, otherwise it returns the generated ID.
func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}

	return id
}

// ConstantGenerator will return the same ID over and over again.
// If no ID is configured, it will return an error.
type ConstantGenerator struct {
	id string
}

// NewConstantGenerator returns a new ConstantGenerator.
func NewConstantGenerator(id string) *ConstantGenerator {
	return &ConstantGenerator{id: id}
}

// New returns the same ID over and over again.
// If no ID is configured, it will return an error.
func (g *ConstantGenerator) New() (string, error) {
	if g.id == "" {
		return "", errors.New("no id is configured")
	}

	return g.id, nil
}

// MustGenerator wraps another generator and delegates the ID generation to it.
// It panics if the delegated generator returns an error.
type MustGenerator struct {
	generator Generator
}

// NewMustGenerator returns a new MustGenerator.
func NewMustGenerator(generator Generator) *MustGenerator {
	return &MustGenerator{generator: generator}
}

// New panics if the underlying generator returns an error, otherwise it returns the generated ID.
func (g *MustGenerator) New() (string, error) {
	if g.generator == nil {
		panic("generator is not configured")
	}

	return Must(g.generator.New()), nil
}
