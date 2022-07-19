package events

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func GetLeadCreatioRequested(lead entity.Lead) ([]byte, error) {
	campaign := &entity.Campaign{
		Medium:    lead.Medium,
		Origin:    lead.Origin,
		AdSetName: lead.AdSetName,
	}
	context := &entity.Context{
		Campaign: campaign,
	}
	properties := entity.Properties{
		Full_Name:       lead.Name,
		Email:           lead.Email,
		Phone:           lead.Phone,
		BusinessHubCode: lead.Hub,
		Description:     lead.Budget,
		ProjectName:     lead.Project,
		ScreenCta:       "LEAD_FORM",
	}
	leadCreationRequested := entity.LeadCreationRequested{
		AnonymousID: uuid.NewString(),
		Event:       "Lead Creation Requested",
		Type:        "Track",
		ReceivedAt:  lead.CreatedAt,
		Context:     *context,
		Properties:  properties,
	}
	request, _ := json.Marshal(leadCreationRequested)
	return request, nil
}

func GetLeadCreatioRequestedEvent(lead entity.Lead) (entity.LeadCreationRequested, error) {
	campaign := &entity.Campaign{
		Medium:    lead.Medium,
		Origin:    lead.Origin,
		AdSet:     lead.AdSetId,
		AdSetName: lead.AdSetName,
		Name:      lead.CampaignName,
		ID:        lead.CampaignId,
		Ad:        lead.AdId,
		AdName:    lead.AdName,
		Source:    lead.Source,
	}
	context := &entity.Context{
		Campaign: campaign,
	}
	LocationOfInterestCodes := []string{}
	if lead.LocationOfInterestCodes != "" {
		LocationOfInterestCodes = []string{lead.LocationOfInterestCodes}
	}

	anonimousID := uuid.NewString()
	if lead.AnonymousID == "" {
		lead.AnonymousID = anonimousID
	}
	properties := entity.Properties{
		Full_Name:                  lead.Name,
		Email:                      lead.Email,
		Phone:                      lead.Phone,
		BusinessHubCode:            lead.Hub,
		ProjectName:                lead.Project,
		ScreenCta:                  "LEAD_FORM",
		LocationOfInterestCodes:    LocationOfInterestCodes,
		MarketingChannel:           lead.MarketingChannel,
		ListingId:                  lead.ProjectId,
		SourcePlatform:             "SPREADSHEET",
		Screen:                     "SPREADSHEET_LEADS",
		MonthlyPaymentBudgetMax:    lead.Monthly_payment_budget,
		BudgetMax:                  lead.BudgetResponse.Max,
		BudgetMin:                  lead.BudgetResponse.Min,
		BudgetCurrency:             lead.BudgetResponse.Currency,
		ExpectedPurchaseTime:       lead.ExpectedPeriod.Time,
		ExpectedPurchaseTimeMax:    lead.ExpectedPeriod.Max,
		ExpectedPurchaseTimeMin:    lead.ExpectedPeriod.Min,
		PreferredPropertyCondition: lead.Preferred_property_condition,
		PurchasePurpose:            lead.Purchase_purpose,
	}
	leadCreationRequested := entity.LeadCreationRequested{
		AnonymousID:       lead.AnonymousID,
		Event:             "Lead Creation Requested",
		Type:              "Track",
		ReceivedAt:        lead.CreatedAt,
		OriginalTimestamp: lead.CreatedAt,
		Context:           *context,
		Properties:        properties,
	}
	return leadCreationRequested, nil
}
