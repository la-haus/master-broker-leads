package serializers

import (
	"log"
	"strconv"
	"time"

	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func GetLead(linesLeads [][]string) []entity.Lead {
	var leads []entity.Lead

	for _, line := range linesLeads {
		name := line[1]
		email := line[2]
		phone := line[3]
		hub := line[4]
		project := line[5]
		medium := line[6]
		status := line[7]
		rowNumber, _ := strconv.Atoi(line[8])
		budget := line[9]
		adSetName := line[10]
		listing_id := line[11]
		if status == "" {
			created_at := time.Now()
			lead := entity.Lead{
				CreatedAt: created_at,
				Name:      name,
				Email:     email,
				Phone:     phone,
				Hub:       hub,
				Project:   project,
				Medium:    medium,
				RowNumber: rowNumber,
				Budget:    budget,
				AdSetName: adSetName,
				Origin:    "SALESFORCE",
				ProjectId: listing_id,
			}
			leads = append(leads, lead)
		}
	}
	return leads
}

func GetLeadsChargeLeads(linesLeads [][]string) []entity.Lead {
	var leads []entity.Lead

	for _, line := range linesLeads {
		rowNumber, err := strconv.Atoi(line[22])
		if err != nil {
			log.Fatal(err)
		}
		created_at := time.Now()
		lead := entity.Lead{
			CreatedAt:                    created_at,
			Name:                         line[1],
			Email:                        line[2],
			Phone:                        line[3],
			Hub:                          line[4],
			LocationOfInterestCodes:      line[5],
			Project:                      line[6],
			ProjectId:                    line[7],
			Source:                       line[8],
			Medium:                       line[9],
			CampaignName:                 line[10],
			CampaignId:                   line[11],
			AdSetName:                    line[12],
			AdSetId:                      line[13],
			AdName:                       line[14],
			AdId:                         line[15],
			Budget:                       line[16],
			Monthly_payment:              line[17],
			Expected_purchase_time:       line[18],
			Purchase_purpose:             line[19],
			Preferred_property_condition: line[20],
			Status:                       line[21],
			RowNumber:                    rowNumber,
			AnonymousID:                  line[23],
		}
		leads = append(leads, lead)
	}
	return leads
}
