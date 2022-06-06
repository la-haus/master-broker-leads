package serializers

import (
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
			}
			leads = append(leads, lead)
		}
	}
	return leads
}
