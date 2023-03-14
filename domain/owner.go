package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type owner struct {
	hash   hash.Hash
	amount uint
	pubKey hash.Hash
}

func createOwner(
	hash hash.Hash,
	amount uint,
	pubKey hash.Hash,
) Owner {
	out := owner{
		hash:   hash,
		amount: amount,
		pubKey: pubKey,
	}

	return &out
}

// Hash returns the hash
func (obj *owner) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *owner) Amount() uint {
	return obj.amount
}

// PubKey returns the pubKey
func (obj *owner) PubKey() hash.Hash {
	return obj.pubKey
}
