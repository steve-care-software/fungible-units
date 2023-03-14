package units

import (
	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the reward application
type Application interface {
	Insert(units units.Units) error
	Retrieve(hash hash.Hash) (units.Unit, error)
	RetrieveByRecipient(recipient hash.Hash) ([]units.Unit, error)
	CalculateReward(pubKeys []hash.Hash, amount uint) (units.Units, error)
}
