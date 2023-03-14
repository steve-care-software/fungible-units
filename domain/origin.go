package domain

import (
	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type origin struct {
	trx  Transaction
	unit units.Unit
}

func createOriginWithTransaction(
	trx Transaction,
) Origin {
	return createOriginInternally(trx, nil)
}

func createOriginWithUnit(
	unit units.Unit,
) Origin {
	return createOriginInternally(nil, unit)
}

func createOriginInternally(
	trx Transaction,
	unit units.Unit,
) Origin {
	out := origin{
		trx:  trx,
		unit: unit,
	}

	return &out
}

// Hash retuns the hash
func (obj *origin) Hash() hash.Hash {
	if obj.IsTransaction() {
		return obj.Transaction().Hash()
	}

	return obj.Unit().Hash()
}

// IsTransaction returns true if there is a transaction, false otherwise
func (obj *origin) IsTransaction() bool {
	return obj.trx != nil
}

// Transaction returns the transaction, if any
func (obj *origin) Transaction() Transaction {
	return obj.trx
}

// IsUnit returns true if there is a unit, false otherwise
func (obj *origin) IsUnit() bool {
	return obj.unit != nil
}

// Unit returns the unit, if any
func (obj *origin) Unit() units.Unit {
	return obj.unit
}
