package events

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func GetLeadCreatioRequested(lead entity.Lead) ([]byte, error) {
	campaign := &entity.Campaign{
		Medium:    lead.Medium,
		Origin:    "SALESFORCE",
		AdSetName: lead.AdSetName,
	}
	context := &entity.Context{
		Campaign: campaign,
	}
	properties := entity.Properties{
		Full_Name:         lead.Name,
		Email:             lead.Email,
		Phone:             lead.Phone,
		Business_hub_code: lead.Hub,
		Description:       lead.Budget,
		ProjectName:       lead.Project,
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
