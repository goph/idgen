package idgen

import (
	"sync"

	"github.com/pkg/errors"
)

// SafeGenerator generates a new ID or returns an error.
// Although receiving errors is highly unlikely, reading from random generator sources could in theory fail.
// Generator implementations should be safe by default.
type SafeGenerator interface {
	Generate() (string, error)
}

// SafeGeneratorFunc generates a new ID or returns an error.
type SafeGeneratorFunc func() (string, error)

func (fn SafeGeneratorFunc) Generate() (string, error) {
	return fn()
}

// Generator generates a new ID and panics in case of any error.
type Generator interface {
	Generate() string
}

// GeneratorFunc generates a new ID and panics in case of any error.
type GeneratorFunc func() string

func (fn GeneratorFunc) Generate() string {
	return fn()
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

// NewGenerator turns a SafeGenerator into a generator that panics when an error occurs.
func NewGenerator(generator SafeGenerator) Generator {
	return &mustGenerator{generator: generator}
}

// mustGenerator wraps another generator and delegates the ID generation to it.
// It panics if the delegated generator returns an error.
type mustGenerator struct {
	generator SafeGenerator
}

// Generate panics if the underlying generator returns an error, otherwise it returns the generated ID.
func (g *mustGenerator) Generate() string {
	if g.generator == nil {
		panic("generator is not configured")
	}

	return Must(g.generator.Generate())
}

// ListGenerator will return an ID from a user-defined list.
// If no ID list is configured or if it reaches the end of the list, it will return an error.
// It is safe to use the ListGenerator in separate goroutines.
type ListGenerator struct {
	ids []string
	cur int
	mu  sync.Mutex
}

// NewListGenerator returns a new ListGenerator.
func NewListGenerator(ids []string) *ListGenerator {
	return &ListGenerator{ids: ids}
}

// Generate returns an ID from a user-defined list.
// If no ID list is configured or if it reaches the end of the list, it will return an error.
// It is safe to call this method in separate goroutines.
func (g *ListGenerator) Generate() (string, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.ids) < g.cur+1 {
		return "", errors.New("no more ids left")
	}

	id := g.ids[g.cur]
	g.cur++

	return id, nil
}
