package entity

import "time"

type Lead struct {
	CreatedAt                    time.Time      `json:"created_at"`
	Name                         string         `json:"name"`
	Email                        string         `json:"email"`
	Phone                        string         `json:"phone,omitempty"`
	Hub                          string         `json:"hub"`
	Project                      string         `json:"project,omitempty"`
	ProjectId                    string         `json:"project_id,omitempty"`
	Medium                       string         `json:"medium,omitempty"`
	Status                       string         `json:"status"`
	Budget                       string         `json:"budget,omitempty"`
	BudgetResponse               Budget         `json:"budget_response,omitempty"`
	AdSetId                      string         `json:"adSet,omitempty"`
	AdSetName                    string         `json:"adSet_name,omitempty"`
	LocationOfInterestCodes      string         `json:"location_of_interest_codes,omitempty"`
	MarketingChannel             string         `json:"marketing_channel,omitempty"`
	CampaignId                   string         `json:"campaign_id,omitempty"`
	CampaignName                 string         `json:"campaign_name,omitempty"`
	AdId                         string         `json:"ad,omitempty"`
	AdName                       string         `json:"ad_name,omitempty"`
	Origin                       string         `json:"origin,omitempty"`
	Screen                       string         `json:"screen,omitempty"`
	SourcePlatform               string         `json:"source_platform,omitempty"`
	Source                       string         `json:"source,omitempty"`
	Monthly_payment              string         `json:"monthly_payment,omitempty"`
	Expected_purchase_time       string         `json:"expected_purchase_time,omitempty"`
	Purchase_purpose             string         `json:"purchase_purpose,omitempty"`
	Preferred_property_condition string         `json:"preferred_property_condition,omitempty"`
	ExpectedPeriod               ExpectedPeriod `json:"expected_period,omitempty"`
	Monthly_payment_budget       float64        `json:"monthly_payment_budget,omitempty"`
	RowNumber                    int
	AnonymousID                  string `json:"anonymous_id,omitempty"`
}
