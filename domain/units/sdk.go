package units

import "github.com/steve-care-software/libs/cryptography/hash"

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