package entity

import "time"

type LeadCreationRequested struct {
	AnonymousID       string     `json:"anonymousId"`
	Event             string     `json:"event"`
	Type              string     `json:"type"`
	ReceivedAt        time.Time  `json:"receivedAt"`
	OriginalTimestamp time.Time  `json:"originalTimestamp"`
	Properties        Properties `json:"properties"`
	Context           Context    `json:"context"`
}
type Context struct {
	Campaign *Campaign `json:"campaign"`
}
type Campaign struct {
	ID               string `json:"id,omitempty"`
	Medium           string `json:"medium,omitempty"`
	Name             string `json:"name,omitempty"`
	Source           string `json:"source,omitempty"`
	Term             string `json:"term,omitempty"`
	Origin           string `json:"origin,omitempty"`
	AdSet            string `json:"adset,omitempty"`
	AdSetName        string `json:"adset_name,omitempty"`
	Ad               string `json:"ad,omitempty"`
	AdName           string `json:"ad_name,omitempty"`
	MarketingChannel string `json:"marketing_channel,omitempty"`
}
type Properties struct {
	Full_Name                  string   `json:"full_name,omitempty"`
	Email                      string   `json:"email"`
	Phone                      string   `json:"phone,omitempty"`
	BusinessHubCode            string   `json:"business_hub_code,omitempty"`
	Description                string   `json:"description,omitempty"`
	ProjectName                string   `json:"project_name,omitempty"`
	ScreenCta                  string   `json:"screen_cta"`
	Screen                     string   `json:"screen,omitempty"`
	LocationOfInterestCodes    []string `json:"location_of_interest_codes,omitempty"`
	MarketingChannel           string   `json:"marketing_channel,omitempty"`
	ListingId                  string   `json:"listing_id,omitempty"`
	SourcePlatform             string   `json:"source_platform,omitempty"`
	ExpectedPurchaseTime       string   `json:"expected_purchase_time,omitempty"`
	ExpectedPurchaseTimeMin    string   `json:"expected_purchase_time_min,omitempty"`
	ExpectedPurchaseTimeMax    string   `json:"expected_purchase_time_max,omitempty"`
	BudgetMax                  float64  `json:"budget_max,omitempty"`
	BudgetMin                  float64  `json:"budget_min,omitempty"`
	BudgetCurrency             string   `json:"budget_currency,omitempty"`
	MonthlyPaymentBudget       float64  `json:"monthly_payment_budget,omitempty"`
	MonthlyPaymentBudgetMax    float64  `json:"monthly_payment_budget_max,omitempty"`
	PreferredPropertyCondition string   `json:"preferred_property_condition,omitempty"`
	PurchasePurpose            string   `json:"purchase_purpose,omitempty"`
}
