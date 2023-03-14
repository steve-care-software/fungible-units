package units

import "github.com/steve-care-software/libs/cryptography/hash"

type units struct {
	hash hash.Hash
	list []Unit
}

func createUnits(
	hash hash.Hash,
	list []Unit,
) Units {
	out := units{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *units) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *units) List() []Unit {
	return obj.list
}
