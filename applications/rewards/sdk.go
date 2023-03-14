package rewards

import (
	"github.com/steve-care-software/fungible-units/domain/units"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// Application represents the reward application
type Application interface {
	Calculate(pubKeys []hash.Hash, amount uint) (units.Units, error)
	Insert(units units.Units) error
}
