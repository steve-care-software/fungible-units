package objects

// Origin represents an origin
type Origin struct {
	Transaction []byte `json:"transaction"`
	Unit        []byte `json:"unit"`
}
