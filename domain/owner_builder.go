package domain

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type ownerBuilder struct {
	hashAdapter hash.Adapter
	amount      uint
	pubKey      hash.Hash
}

func createOwnerBuilder(
	hashAdapter hash.Adapter,
) OwnerBuilder {
	out := ownerBuilder{
		hashAdapter: hashAdapter,
		amount:      0,
		pubKey:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *ownerBuilder) Create() OwnerBuilder {
	return createOwnerBuilder(app.hashAdapter)
}

// WithAmount adds an amount to the builder
func (app *ownerBuilder) WithAmount(amount uint) OwnerBuilder {
	app.amount = amount
	return app
}

// WithPubKey adds a pubKey to the builder
func (app *ownerBuilder) WithPubKey(pubKey hash.Hash) OwnerBuilder {
	app.pubKey = pubKey
	return app
}

// Now builds a new Owner instance
func (app *ownerBuilder) Now() (Owner, error) {
	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build an Owner instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the pubKey is mandatory in order to build an Owner instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", app.amount)),
		app.pubKey.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createOwner(*pHash, app.amount, app.pubKey), nil
}
