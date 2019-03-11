package uuidgen

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"
)

// nolint: gochecknoglobals
var namespaceDNS = uuid.Must(uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))

func TestNewGenerator(t *testing.T) {
	type gen interface {
		Generate() (string, error)
	}
	tests := []func() gen{
		func() gen { return NewV1Generator() },
		func() gen { return NewV2Generator(byte(1)) },
		func() gen { return NewV3Generator(namespaceDNS, "v3") },
		func() gen { return NewV4Generator() },
		func() gen { return NewV5Generator(namespaceDNS, "v5") },
	}

	for i, test := range tests {
		i, test := i, test

		t.Run(fmt.Sprintf("v%d", i+1), func(t *testing.T) {
			generator := test()

			id, err := generator.Generate()
			if err != nil {
				t.Fatal(err)
			}

			if id == "" {
				t.Errorf("the generator is expected to work with default options")
			}
		})
	}
}

type uuidGen struct {
	uuid string
}

func (g *uuidGen) NewV1() (uuid.UUID, error) {
	return uuid.FromString(g.uuid)
}

func (g *uuidGen) NewV2(domain byte) (uuid.UUID, error) {
	return uuid.FromString(g.uuid)
}

func (g *uuidGen) NewV3(ns uuid.UUID, name string) uuid.UUID {
	return uuid.FromStringOrNil(g.uuid)
}

func (g *uuidGen) NewV4() (uuid.UUID, error) {
	return uuid.FromString(g.uuid)
}

func (g *uuidGen) NewV5(ns uuid.UUID, name string) uuid.UUID {
	return uuid.FromStringOrNil(g.uuid)
}

func TestGenerator(t *testing.T) {
	type gen interface {
		Generate() (string, error)
	}
	tests := []struct {
		generator func(g *uuidGen) gen
		uuid      string
	}{
		{
			func(g *uuidGen) gen { return NewV1GeneratorWithGen(g) },
			"7ff4ebf2-439c-11e9-b210-d663bd873d93",
		},
		{
			func(g *uuidGen) gen { return NewV2GeneratorWithGen(byte(1), g) },
			"7ff4ebf2-439c-21e9-b210-d663bd873d93",
		},
		{
			func(g *uuidGen) gen { return NewV3GeneratorWithGen(namespaceDNS, "v3", g) },
			"a3bb189e-8bf9-3888-9912-ace4e6543002",
		},
		{
			func(g *uuidGen) gen { return NewV4GeneratorWithGen(g) },
			"27d3a370-1c18-46b7-947f-8a81b3f867e0",
		},
		{
			func(g *uuidGen) gen { return NewV5GeneratorWithGen(namespaceDNS, "v3", g) },
			"a6edc906-2f9f-5fb2-a373-efac406f0ef2",
		},
	}

	for i, test := range tests {
		i, test := i, test

		t.Run(fmt.Sprintf("v%d", i+1), func(t *testing.T) {
			generator := test.generator(&uuidGen{test.uuid})

			id, err := generator.Generate()
			if err != nil {
				t.Fatal(err)
			}

			if got, want := id, test.uuid; got != want {
				t.Errorf("the generated id does not match the expected\nexpected: %s\nactual:   %s", want, got)
			}
		})
	}
}
