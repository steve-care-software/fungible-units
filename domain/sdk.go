package domain

import (
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewOwnersBuilder creates a new owners builder
func NewOwnersBuilder() OwnersBuilder {
	hashAdapter := hash.NewAdapter()
	return createOwnersBuilder(hashAdapter)
}

// NewOwnerBuilder creates a new owner instance
func NewOwnerBuilder() OwnerBuilder {
	hashAdapter := hash.NewAdapter()
	return createOwnerBuilder(hashAdapter)
}

// NewOriginsBuilder creates a new origins builder
func NewOriginsBuilder() OriginsBuilder {
	hashAdapter := hash.NewAdapter()
	return createOriginsBuilder(hashAdapter)
}

// NewOriginBuilder creates a new origin builder
func NewOriginBuilder() OriginBuilder {
	return createOriginBuilder()
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithGenesis(genesis hash.Hash) Builder
	WithOrigins(origins Origins) Builder
	WithOwners(owners Owners) Builder
	ExecutedOn(executedOn time.Time) Builder
	FreezeUntil(freezeUntil time.Time) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Genesis() hash.Hash
	Origins() Origins
	Owners() Owners
	HasExecuteOn() bool
	ExecuteOn() time.Time
	HasFreezeUntil() bool
	FreezeUntil() time.Time
}

// OwnersBuilder represents an owners builder
type OwnersBuilder interface {
	Create() OwnersBuilder
	WithList(list []Owner) OwnersBuilder
	Now() (Owners, error)
}

// Owners represents owners
type Owners interface {
	Hash() hash.Hash
	List() []Owner
}

// OwnerBuilder represents an owner builder
type OwnerBuilder interface {
	Create() OwnerBuilder
	WithAmount(amount uint) OwnerBuilder
	WithPubKey(pubKey hash.Hash) OwnerBuilder
	Now() (Owner, error)
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Amount() uint
	PubKey() hash.Hash
}

// OriginsBuilder represents an origins builder
type OriginsBuilder interface {
	Create() OriginsBuilder
	WithList(list []Origin) OriginsBuilder
	Now() (Origins, error)
}

// Origins represents origins
type Origins interface {
	Hash() hash.Hash
	List() []Origin
}

// OriginBuilder represents an origin builder
type OriginBuilder interface {
	Create() OriginBuilder
	WithTransaction(trx Transaction) OriginBuilder
	WithUnit(unit units.Unit) OriginBuilder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	IsTransaction() bool
	Transaction() Transaction
	IsUnit() bool
	Unit() units.Unit
}
