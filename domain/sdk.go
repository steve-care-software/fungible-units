package domain

import (
	"time"

	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

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
	ExecuteOn(executeOn time.Time) Builder
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
	ExecuteOn() *time.Time
	HasFreezeUntil() bool
	FreezeUntil() *time.Time
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithContext(context uint) RepositoryBuilder
	WithKind(kind uint) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a transaction repository
type Repository interface {
	Retrieve(hash hash.Hash) (Transaction, error)
	RetrieveBySender(sender hash.Hash) ([]Transaction, error)
	RetrieveByRecipient(recipient hash.Hash) ([]Transaction, error)
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Create() ServiceBuilder
	WithContext(context uint) ServiceBuilder
	WithKind(kind uint) ServiceBuilder
	Now() (Service, error)
}

// Service represents a transaction service
type Service interface {
	Insert(trx Transaction) error
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

// OwnerRepositoryBuilder represents an owner repository builder
type OwnerRepositoryBuilder interface {
	Create() OwnerRepositoryBuilder
	WithContext(context uint) OwnerRepositoryBuilder
	WithKind(kind uint) OwnerRepositoryBuilder
	Now() (OwnerRepository, error)
}

// OwnerRepository represents an owner repository
type OwnerRepository interface {
	Retrieve(hash hash.Hash) (Owner, error)
}

// OwnerServiceBuilder represents an owner service builder
type OwnerServiceBuilder interface {
	Create() OwnerServiceBuilder
	WithContext(context uint) OwnerServiceBuilder
	WithKind(kind uint) OwnerServiceBuilder
	Now() (OwnerService, error)
}

// OwnerService represents an owner service
type OwnerService interface {
	Insert(trx Owner) error
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

// OriginRepositoryBuilder represents an origin repository builder
type OriginRepositoryBuilder interface {
	Create() OriginRepositoryBuilder
	WithContext(context uint) OriginRepositoryBuilder
	WithKind(kind uint) OriginRepositoryBuilder
	Now() (OriginRepository, error)
}

// OriginRepository represents an origin repository
type OriginRepository interface {
	Retrieve(hash hash.Hash) (Origin, error)
}

// OriginServiceBuilder represents an origin service builder
type OriginServiceBuilder interface {
	Create() OriginServiceBuilder
	WithContext(context uint) OriginServiceBuilder
	WithKind(kind uint) OriginServiceBuilder
	Now() (OriginService, error)
}

// OriginService represents an origin service
type OriginService interface {
	Insert(trx Origin) error
}
