package domain

import (
	"errors"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type ownersBuilder struct {
	hashAdapter hash.Adapter
	list        []Owner
}

func createOwnersBuilder(
	hashAdapter hash.Adapter,
) OwnersBuilder {
	out := ownersBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *ownersBuilder) Create() OwnersBuilder {
	return createOwnersBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *ownersBuilder) WithList(list []Owner) OwnersBuilder {
	app.list = list
	return app
}

// Now builds a new Owners instance
func (app *ownersBuilder) Now() (Owners, error) {
	if app.list == nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Owner in order to build a Owners instance")
	}

	data := [][]byte{}
	for _, oneOwner := range app.list {
		data = append(data, oneOwner.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createOwners(*pHash, app.list), nil
}
