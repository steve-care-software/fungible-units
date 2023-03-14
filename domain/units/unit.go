package units

import "github.com/steve-care-software/libs/cryptography/hash"

type unit struct {
	hash    hash.Hash
	genesis hash.Hash
	pubKey  hash.Hash
	amount  uint
}

func createUnit(
	hash hash.Hash,
	genesis hash.Hash,
	pubKey hash.Hash,
	amount uint,
) Unit {
	out := unit{
		hash:    hash,
		genesis: genesis,
		pubKey:  pubKey,
		amount:  amount,
	}

	return &out
}

// Hash returns the hash
func (obj *unit) Hash() hash.Hash {
	return obj.hash
}

// Genesis returns the genesis hash
func (obj *unit) Genesis() hash.Hash {
	return obj.genesis
}

// PubKey returns the pubKey
func (obj *unit) PubKey() hash.Hash {
	return obj.pubKey
}

// Amount returns the amount
func (obj *unit) Amount() uint {
	return obj.amount
}
