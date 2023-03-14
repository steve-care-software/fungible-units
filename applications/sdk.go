package applications

import (
	"github.com/steve-care-software/fungible-units/applications/transactions"
	"github.com/steve-care-software/fungible-units/applications/units"
)

// Application represents an application
type Application interface {
	Transaction() transactions.Application
	Unit() units.Application
}
