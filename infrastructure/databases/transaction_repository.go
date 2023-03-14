package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	transactions "github.com/steve-care-software/fungible-units/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionRepository struct {
	hashAdapter      hash.Adapter
	database         database_application.Application
	ownerRepository  transactions.OwnerRepository
	originRepository transactions.OriginRepository
	builder          transactions.Builder
	context          uint
	kind             uint
}
