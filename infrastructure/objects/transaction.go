package objects

import "time"

// Transaction represents a transaction
type Transaction struct {
	Genesis     []byte     `json:"genesis"`
	Origins     [][]byte   `json:"origins"`
	Owners      [][]byte   `json:"owners"`
	ExecuteOn   *time.Time `json:"execute_on"`
	FreezeUntil *time.Time `json:"freeze_until"`
}
