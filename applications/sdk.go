package applications

import (
	transactions "github.com/steve-care-software/fungible-units/domain"
	"github.com/steve-care-software/libs/cryptography/hash"
)

// EnterOnProcessTrxFn represents the enter's onProcess Trx func
type EnterOnProcessTrxFn func(trx transactions.Transaction) error

// ExitOnProcessTrxFn represents the exit's onProcess Trx func
type ExitOnProcessTrxFn func(trx transactions.Transaction) error

// Application represents a transaction application
type Application interface {
	Transact(trx transactions.Transaction) error
	Retrieve(hash hash.Hash) (transactions.Transaction, error)
	RetrieveBySender(sender hash.Hash) ([]transactions.Transaction, error)
	RetrieveByRecipient(recipient hash.Hash) ([]transactions.Transaction, error)
}
