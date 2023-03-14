package objects

// Unit represents a unit
type Unit struct {
	Genesis []byte `json:"genesis"`
	PubKey  []byte `json:"pubkey"`
	Amount  uint   `json:"amount"`
}
