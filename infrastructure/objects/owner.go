package objects

// Owner represents an owner
type Owner struct {
	Amount uint   `json:"amount"`
	PubKey []byte `json:"pubkey"`
}
