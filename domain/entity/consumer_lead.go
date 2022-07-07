package entity

import "time"

type LeadCreationRequested struct {
	AnonymousID string     `json:"anonymousId"`
	Event       string     `json:"event"`
	Type        string     `json:"type"`
	ReceivedAt  time.Time  `json:"receivedAt"`
	Properties  Properties `json:"properties"`
	Context     Context    `json:"context"`
}
type Context struct {
	Campaign *Campaign `json:"campaign"`
}
type Campaign struct {
	Medium    string `json:"medium,omitempty"`
	Origin    string `json:"origin,omitempty"`
	AdSetName string `json:"adset_name,omitempty"`
}
type Properties struct {
	Full_Name         string `json:"full_name,omitempty"`
	Email             string `json:"email"`
	Phone             string `json:"phone,omitempty"`
	Business_hub_code string `json:"business_hub_code"`
	Description       string `json:"description,omitempty"`
	ProjectName       string `json:"project_name,omitempty"`
	Screen_cta        string `json:"screen_cta"`
}
