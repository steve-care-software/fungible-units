package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type originsBuilder struct {
	hashAdapter hash.Adapter
	list        []Origin
}

func createOriginsBuilder(
	hashAdapter hash.Adapter,
) OriginsBuilder {
	out := originsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *originsBuilder) Create() OriginsBuilder {
	return createOriginsBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *originsBuilder) WithList(list []Origin) OriginsBuilder {
	app.list = list
	return app
}

// Now builds a new Origins instance
func (app *originsBuilder) Now() (Origins, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Origin in order to build a Origins instance")
	}

	data := [][]byte{}
	for _, oneOrigin := range app.list {
		data = append(data, oneOrigin.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createOrigins(*pHash, app.list), nil
}
