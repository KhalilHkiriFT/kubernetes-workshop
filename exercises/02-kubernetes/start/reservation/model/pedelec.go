package model

type Pedelec struct {
	ID            string  `json:"id"`
	Charge        int     `json:"charge"`
	IsAvailable   bool    `json:"isAvailable"`
	WorkingRadius float64 `json:"workingRadius"`
}
