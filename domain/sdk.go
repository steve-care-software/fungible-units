package domain

import (
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

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
	WithOrigin(origin hash.Hash) OriginBuilder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	IsTransaction() bool
	Transaction() Transaction
	IsOrigin() bool
	Origin() hash.Hash
}
