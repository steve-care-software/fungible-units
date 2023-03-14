package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

type transaction struct {
	hash         hash.Hash
	genesis      hash.Hash
	origins      Origins
	owners       Owners
	pExecuteOn   *time.Time
	pFreezeUntil *time.Time
}

func createTransaction(
	hash hash.Hash,
	genesis hash.Hash,
	origins Origins,
	owners Owners,
) Transaction {
	return createTransactionInternally(hash, genesis, origins, owners, nil, nil)
}

func createTransactionWithExecuteOn(
	hash hash.Hash,
	genesis hash.Hash,
	origins Origins,
	owners Owners,
	pExecuteOn *time.Time,
) Transaction {
	return createTransactionInternally(hash, genesis, origins, owners, pExecuteOn, nil)
}

func createTransactionWithFreezeUntil(
	hash hash.Hash,
	genesis hash.Hash,
	origins Origins,
	owners Owners,
	pFreezeUntil *time.Time,
) Transaction {
	return createTransactionInternally(hash, genesis, origins, owners, nil, pFreezeUntil)
}

func createTransactionWithExecuteOnAndFreezeUntil(
	hash hash.Hash,
	genesis hash.Hash,
	origins Origins,
	owners Owners,
	pExecuteOn *time.Time,
	pFreezeUntil *time.Time,
) Transaction {
	return createTransactionInternally(hash, genesis, origins, owners, pExecuteOn, pFreezeUntil)
}

func createTransactionInternally(
	hash hash.Hash,
	genesis hash.Hash,
	origins Origins,
	owners Owners,
	pExecuteOn *time.Time,
	pFreezeUntil *time.Time,
) Transaction {
	out := transaction{
		hash:         hash,
		genesis:      genesis,
		origins:      origins,
		owners:       owners,
		pExecuteOn:   pExecuteOn,
		pFreezeUntil: pFreezeUntil,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Genesis returns the genesis
func (obj *transaction) Genesis() hash.Hash {
	return obj.genesis
}

// Origins returns the origins
func (obj *transaction) Origins() Origins {
	return obj.origins
}

// Owners returns the owners
func (obj *transaction) Owners() Owners {
	return obj.owners
}

// HasExecuteOn returns true if there is an pExecuteOn, false otherwise
func (obj *transaction) HasExecuteOn() bool {
	return obj.pExecuteOn != nil
}

// ExecuteOn returns the execution time, if any
func (obj *transaction) ExecuteOn() *time.Time {
	return obj.pExecuteOn
}

// HasFreezeUntil returns true if there is a pFreezeUntil, false otherwise
func (obj *transaction) HasFreezeUntil() bool {
	return obj.pFreezeUntil != nil
}

// FreezeUntil returns the pFreezeUntil, if any
func (obj *transaction) FreezeUntil() *time.Time {
	return obj.pFreezeUntil
}
