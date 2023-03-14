package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	genesis      hash.Hash
	origins      Origins
	owners       Owners
	pExecuteOn   *time.Time
	pFreezeUntil *time.Time
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		genesis:      nil,
		origins:      nil,
		owners:       nil,
		pExecuteOn:   nil,
		pFreezeUntil: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithGenesis adds a genesis to the builder
func (app *builder) WithGenesis(genesis hash.Hash) Builder {
	app.genesis = genesis
	return app
}

// WithOrigins adds an origins to the builder
func (app *builder) WithOrigins(origins Origins) Builder {
	app.origins = origins
	return app
}

// WithOwners adds an owners to the builder
func (app *builder) WithOwners(owners Owners) Builder {
	app.owners = owners
	return app
}

// ExecuteOn adds an executeOn to the builder
func (app *builder) ExecuteOn(executeOn time.Time) Builder {
	app.pExecuteOn = &executeOn
	return app
}

// FreezeUntil adds a freezeUntil to the builder
func (app *builder) FreezeUntil(freezeUntil time.Time) Builder {
	app.pFreezeUntil = &freezeUntil
	return app
}

// Now builds a new Transaction instance
func (app *builder) Now() (Transaction, error) {
	if app.genesis == nil {
		return nil, errors.New("the genesis hash is mandatory in order to build a Transaction instance")
	}

	if app.origins == nil {
		return nil, errors.New("the origins is mandatory in order to build a Transaction instance")
	}

	if app.owners == nil {
		return nil, errors.New("the owners is mandatory in order to build a Transaction instance")
	}

	data := [][]byte{
		app.genesis.Bytes(),
		app.origins.Hash().Bytes(),
		app.owners.Hash().Bytes(),
	}

	if app.pExecuteOn != nil {
		data = append(data, []byte(fmt.Sprintf("%d", app.pExecuteOn.Unix())))
	}

	if app.pFreezeUntil == nil {
		data = append(data, []byte(fmt.Sprintf("%d", app.pFreezeUntil.Unix())))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.pExecuteOn != nil && app.pFreezeUntil != nil {
		return createTransactionWithExecuteOnAndFreezeUntil(*pHash, app.genesis, app.origins, app.owners, app.pExecuteOn, app.pFreezeUntil), nil
	}

	if app.pExecuteOn != nil {
		return createTransactionWithExecuteOn(*pHash, app.genesis, app.origins, app.owners, app.pExecuteOn), nil
	}

	if app.pFreezeUntil == nil {
		return createTransactionWithFreezeUntil(*pHash, app.genesis, app.origins, app.owners, app.pFreezeUntil), nil
	}

	return createTransaction(*pHash, app.genesis, app.origins, app.owners), nil
}
