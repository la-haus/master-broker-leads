package entity

import "time"

type Lead struct {
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Hub       string    `json:"hub"`
	Project   string    `json:"project,omitempty"`
	Medium    string    `json:"medium,omitempty"`
	Status    string    `json:"status"`
	Budget    string    `json:"budget,omitempty"`
	AdSetName string    `json:"adSet_name,omitempty"`
	RowNumber int
}
