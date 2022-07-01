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
	Medium    string `json:"medium"`
	Origin    string `json:"origin"`
	AdSetName string `json:"adset_name"`
}
type Properties struct {
	Full_Name         string `json:"full_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Business_hub_code string `json:"business_hub_code"`
	Description       string `json:"description"`
	ProjectName       string `json:"project_name"`
	Screen_cta        string `json:"screen_cta"`
}
