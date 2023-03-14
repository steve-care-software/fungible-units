package units

import "github.com/steve-care-software/libs/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewUnitBuilder creates a new unit builder
func NewUnitBuilder() UnitBuilder {
	hashAdapter := hash.NewAdapter()
	return createUnitBuilder(hashAdapter)
}

// Builder represents the units builder
type Builder interface {
	Create() Builder
	WithList(list []Unit) Builder
	Now() (Units, error)
}

// Units represents units
type Units interface {
	Hash() hash.Hash
	List() []Unit
}

// UnitBuilder represents the unit builder
type UnitBuilder interface {
	Create() UnitBuilder
	WithGenesis(genesis hash.Hash) UnitBuilder
	WithPubKey(pubKey hash.Hash) UnitBuilder
	WithAmount(amount uint) UnitBuilder
	Now() (Unit, error)
}

// Unit represents an origin unit
type Unit interface {
	Hash() hash.Hash
	Genesis() hash.Hash
	PubKey() hash.Hash
	Amount() uint
}

// UnitRepositoryBuilder represents an unit repository builder
type UnitRepositoryBuilder interface {
	Create() UnitRepositoryBuilder
	WithContext(context uint) UnitRepositoryBuilder
	WithKind(kind uint) UnitRepositoryBuilder
	Now() (UnitRepository, error)
}

// UnitRepository represents an unit repository
type UnitRepository interface {
	Retrieve(hash hash.Hash) (Unit, error)
}

// UnitServiceBuilder represents an unit service builder
type UnitServiceBuilder interface {
	Create() UnitServiceBuilder
	WithContext(context uint) UnitServiceBuilder
	WithKind(kind uint) UnitServiceBuilder
	Now() (UnitService, error)
}

// UnitService represents an unit service
type UnitService interface {
	Insert(trx Unit) error
}
