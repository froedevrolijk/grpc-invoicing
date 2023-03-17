package models

type Pricing struct {
	Id     int64   `json:"id" gorm:"primaryKey"`
	Rate   float64 `json:"rate"`
	Charge bool    `json:"charge"`
	Amount float64 `json:"amount"`
}
