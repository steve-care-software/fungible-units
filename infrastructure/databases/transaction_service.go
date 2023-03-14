package databases

import (
	database_application "github.com/steve-care-software/databases/applications"
	transactions "github.com/steve-care-software/fungible-units/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

type transactionService struct {
	hashAdapter      hash.Adapter
	database         database_application.Application
	repository       transactions.Repository
	ownerRepository  transactions.OwnerRepository
	originRepository transactions.OriginRepository
	context          uint
	kind             uint
}
