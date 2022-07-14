package entity

type Budget struct {
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Currency string  `json:"currency"`
}
