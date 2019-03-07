package idgen

import (
	"errors"
)

// SafeGenerator generates a new ID or returns an error.
// Although receiving errors is highly unlikely, reading from random generator sources could in theory fail.
// Generator implementations should be safe by default.
type SafeGenerator interface {
	Generate() (string, error)
}

// Generator generates a new ID and panics in case of any error.
type Generator interface {
	Generate() string
}

// Must panics if a safe generator function returns with an error, otherwise it returns the generated ID.
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

// Generate returns the same ID over and over again.
// If no ID is configured, it will return an error.
func (g *ConstantGenerator) Generate() (string, error) {
	if g.id == "" {
		return "", errors.New("no id is configured")
	}

	return g.id, nil
}

// MustGenerator wraps another generator and delegates the ID generation to it.
// It panics if the delegated generator returns an error.
type MustGenerator struct {
	generator SafeGenerator
}

// NewMustGenerator returns a new MustGenerator.
func NewMustGenerator(generator SafeGenerator) *MustGenerator {
	return &MustGenerator{generator: generator}
}

// Generate panics if the underlying generator returns an error, otherwise it returns the generated ID.
func (g *MustGenerator) Generate() string {
	if g.generator == nil {
		panic("generator is not configured")
	}

	return Must(g.generator.Generate())
}
