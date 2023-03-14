package units

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Unit
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Unit) Builder {
	app.list = list
	return app
}

// Now builds a new Units instance
func (app *builder) Now() (Units, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Unit in order to build a Units instance")
	}

	data := [][]byte{}
	for _, oneUnit := range app.list {
		data = append(data, oneUnit.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createUnits(*pHash, app.list), nil
}
