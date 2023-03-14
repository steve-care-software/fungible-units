package domain

import (
	"errors"

	"github.com/steve-care-software/fungible-units/domain/units"
)

type originBuilder struct {
	trx  Transaction
	unit units.Unit
}

func createOriginBuilder() OriginBuilder {
	out := originBuilder{
		trx:  nil,
		unit: nil,
	}

	return &out
}

// Create initializes the builder
func (app *originBuilder) Create() OriginBuilder {
	return createOriginBuilder()
}

// WithTransaction adds a transaction to the builder
func (app *originBuilder) WithTransaction(trx Transaction) OriginBuilder {
	app.trx = trx
	return app
}

// WithUnit adds a unit to the builder
func (app *originBuilder) WithUnit(unit units.Unit) OriginBuilder {
	app.unit = unit
	return app
}

// Now builds a new Origin insatance
func (app *originBuilder) Now() (Origin, error) {
	if app.trx != nil {
		return createOriginWithTransaction(app.trx), nil
	}

	if app.unit != nil {
		return createOriginWithUnit(app.unit), nil
	}

	return nil, errors.New("the Origin is invalid")
}
