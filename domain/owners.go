package domain

import "github.com/steve-care-software/libs/cryptography/hash"

type owners struct {
	hash hash.Hash
	list []Owner
}

func createOwners(
	hash hash.Hash,
	list []Owner,
) Owners {
	out := owners{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *owners) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *owners) List() []Owner {
	return obj.list
}
