package charge_leads_job

import (
	"fmt"
	"log"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	"github.com/la-haus/master-brokers-charge-leads/domain/events"
	"github.com/la-haus/master-brokers-charge-leads/domain/request_lead"
	"github.com/la-haus/master-brokers-charge-leads/domain/serializers"
	"github.com/la-haus/master-brokers-charge-leads/domain/validator"
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
	srv := google_func.Conn(charge.config)
	spreadsheetId := "1bS_OYWaOApCEodQBqT6kosMKKs7llNt8hAqOdg7RKm8"
	data, err := google_func.ReadSpreadSheet(srv, spreadsheetId)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	leads := serializers.GetLead(data)
	for _, lead := range leads {
		lead.Phone = validator.ValidatePhone(lead.Phone, lead.Hub[:2], charge.config)
		event, _ := events.GetLeadCreatioRequestedEvent(lead)
		SegmentClient := request_lead.NewSegmentClient(charge.config)
		defer SegmentClient.Client.Close()
		err := SegmentClient.SendTrackLead(event)
		if err != nil {
			fmt.Println("Error send event: ", err)
			google_func.WriteSpreadSheet(srv, spreadsheetId, lead, false)
			continue
		}
		google_func.WriteSpreadSheet(srv, spreadsheetId, lead, true)
	}
}
