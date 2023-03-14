package units

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type unitBuilder struct {
	hashAdapter hash.Adapter
	genesis     hash.Hash
	pubKey      hash.Hash
	amount      uint
}

func createUnitBuilder(
	hashAdapter hash.Adapter,
) UnitBuilder {
	out := unitBuilder{
		hashAdapter: hashAdapter,
		genesis:     nil,
		pubKey:      nil,
		amount:      0,
	}

	return &out
}

// Create initializes the builder
func (app *unitBuilder) Create() UnitBuilder {
	return createUnitBuilder(app.hashAdapter)
}

// WithGenesis adds a genesis to the builder
func (app *unitBuilder) WithGenesis(genesis hash.Hash) UnitBuilder {
	app.genesis = genesis
	return app
}

// WithPubKey adds a pubKey to the builder
func (app *unitBuilder) WithPubKey(pubKey hash.Hash) UnitBuilder {
	app.pubKey = pubKey
	return app
}

// WithAmount adds an amount to the builder
func (app *unitBuilder) WithAmount(amount uint) UnitBuilder {
	app.amount = amount
	return app
}

// Now builds a new Unit instance
func (app *unitBuilder) Now() (Unit, error) {
	if app.genesis == nil {
		return nil, errors.New("the genesis hash is mandatory in order to build a Unit instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the pubKey hash is mandatory in order to build a Unit instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Unit instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.genesis.Bytes(),
		app.pubKey.Bytes(),
		[]byte(fmt.Sprintf("%d", app.amount)),
	})

	if err != nil {
		return nil, err
	}

	return createUnit(
		*pHash,
		app.genesis,
		app.pubKey,
		app.amount,
	), nil
}
