package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type origins struct {
	hash hash.Hash
	list []Origin
}

func createOrigins(
	hash hash.Hash,
	list []Origin,
) Origins {
	out := origins{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *origins) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *origins) List() []Origin {
	return obj.list
}
