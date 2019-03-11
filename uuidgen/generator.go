package uuidgen

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// V1Generator generates a UUID V1.
type V1Generator struct {
	gen uuid.Generator
}

// NewV1Generator returns a new UUID V1 generator.
func NewV1Generator() *V1Generator {
	return NewV1GeneratorWithGen(uuid.NewGen())
}

// NewV1GeneratorWithGen returns a new UUID V1 generator.
func NewV1GeneratorWithGen(gen uuid.Generator) *V1Generator {
	return &V1Generator{
		gen: gen,
	}
}

// Generate generates a UUID V1.
func (g *V1Generator) Generate() (string, error) {
	id, err := g.gen.NewV1()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}

// V2Generator generates a UUID V2.
type V2Generator struct {
	domain byte

	gen uuid.Generator
}

// NewV2Generator returns a new UUID V2 generator.
func NewV2Generator(domain byte) *V2Generator {
	return NewV2GeneratorWithGen(domain, uuid.NewGen())
}

// NewV2GeneratorWithGen returns a new UUID V2 generator.
func NewV2GeneratorWithGen(domain byte, gen uuid.Generator) *V2Generator {
	return &V2Generator{
		domain: domain,

		gen: gen,
	}
}

// Generate generates a UUID V2.
func (g *V2Generator) Generate() (string, error) {
	id, err := g.gen.NewV2(g.domain)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}

// V3Generator generates a UUID V3.
type V3Generator struct {
	ns   uuid.UUID
	name string

	gen uuid.Generator
}

// NewV3Generator returns a new UUID V3 generator.
func NewV3Generator(ns uuid.UUID, name string) *V3Generator {
	return NewV3GeneratorWithGen(ns, name, uuid.NewGen())
}

// NewV3GeneratorWithGen returns a new UUID V3 generator.
func NewV3GeneratorWithGen(ns uuid.UUID, name string, gen uuid.Generator) *V3Generator {
	return &V3Generator{
		ns:   ns,
		name: name,

		gen: gen,
	}
}

// Generate generates a UUID V3.
func (g *V3Generator) Generate() (string, error) {
	return g.gen.NewV3(g.ns, g.name).String(), nil
}

// V4Generator generates a UUID V4.
type V4Generator struct {
	gen uuid.Generator
}

// NewV4Generator returns a new UUID V4 generator.
func NewV4Generator() *V4Generator {
	return NewV4GeneratorWithGen(uuid.NewGen())
}

// NewV4GeneratorWithGen returns a new UUID V4 generator.
func NewV4GeneratorWithGen(gen uuid.Generator) *V4Generator {
	return &V4Generator{
		gen: gen,
	}
}

// Generate generates a UUID V4.
func (g *V4Generator) Generate() (string, error) {
	id, err := g.gen.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate ID")
	}

	return id.String(), nil
}

// V5Generator generates a UUID V5.
type V5Generator struct {
	ns   uuid.UUID
	name string

	gen uuid.Generator
}

// NewV5Generator returns a new UUID V5 generator.
func NewV5Generator(ns uuid.UUID, name string) *V5Generator {
	return NewV5GeneratorWithGen(ns, name, uuid.NewGen())
}

// NewV5GeneratorWithGen returns a new UUID V5 generator.
func NewV5GeneratorWithGen(ns uuid.UUID, name string, gen uuid.Generator) *V5Generator {
	return &V5Generator{
		ns:   ns,
		name: name,

		gen: gen,
	}
}

// Generate generates a UUID V5.
func (g *V5Generator) Generate() (string, error) {
	return g.gen.NewV5(g.ns, g.name).String(), nil
}
