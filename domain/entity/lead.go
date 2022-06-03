package entity

import "time"

type Lead struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Hub       string    `json:"hub"`
	Project   string    `json:"project"`
	Medium    string    `json:"medium"`
	Status    string    `json:"status"`
	Budget    string    `json:"budget"`
	AdSetName string    `json:"adSet_name"`
	RowNumber int
}
