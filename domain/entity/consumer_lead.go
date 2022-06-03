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
	ID        string `json:"id"`
	Medium    string `json:"medium"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Term      string `json:"term"`
	Origin    string `json:"origin"`
	AdSet     string `json:"adset"`
	AdSetName string `json:"adset_name"`
	Ad        string `json:"ad"`
	AdName    string `json:"ad_name"`
}
type Properties struct {
	Full_Name         string `json:"full_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Business_hub_code string `json:"business_hub_code"`
	Description       string `json:"description"`
	ProjectName       string `json:"project_name"`
}
