package charge_leads_job

import (
	"log"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	"github.com/la-haus/master-brokers-charge-leads/domain/events"
	"github.com/la-haus/master-brokers-charge-leads/domain/request_lead"
	"github.com/la-haus/master-brokers-charge-leads/domain/serializers"
	"github.com/la-haus/master-brokers-charge-leads/util/google_func"
)

type ChargeLeadsUseCase interface {
	ChargeLeads()
}

func NewChargeLeadsUseCase(config *configs.Config) ChargeLeadsUseCase {
	return &chargeLeadsUseCase{
		config: config,
	}
}

type chargeLeadsUseCase struct {
	config *configs.Config
}

func (charge *chargeLeadsUseCase) ChargeLeads() {
	srv := google_func.ConfigWrite()
	spreadsheetId := "1bS_OYWaOApCEodQBqT6kosMKKs7llNt8hAqOdg7RKm8"
	data, err := google_func.ReadSpreadSheet(srv, spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	leads := serializers.GetLead(data)
	for _, lead := range leads {
		event, _ := events.GetLeadCreatioRequested(lead)
		err := request_lead.SendEvent(event, charge.config)
		if err != nil {
			google_func.WriteSpreadSheet(srv, spreadsheetId, lead, false)
			continue
		}
		google_func.WriteSpreadSheet(srv, spreadsheetId, lead, true)
	}
}
