package google_func

import (
	"fmt"
	"log"
	"strconv"

	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
	"google.golang.org/api/sheets/v4"
)

type ServiceClient struct {
	Srv *sheets.Service
}

func NewServiceClient(srv *sheets.Service) *ServiceClient {
	return &ServiceClient{
		Srv: srv,
	}
}

func (sc *ServiceClient) ReadSpreadSheet(spreadsheetId string) ([][]string, error) {
	var tableLeads [][]string
	readRange := "leads!K2:V"
	resp, err := sc.Srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	if len(resp.Values) == 0 {
		return tableLeads, nil
	} else {
		for number, row := range resp.Values {
			positionInSpreadsheet := number + 2
			dataIsComplete := fmt.Sprintf("%s", row[0])
			if dataIsComplete == "DATOS COMPLETOS" {
				name := fmt.Sprintf("%s", row[1])
				email := fmt.Sprintf("%s", row[2])
				phone := fmt.Sprintf("%s", row[3])
				hub := fmt.Sprintf("%s", row[4])
				project := fmt.Sprintf("%s", row[5])
				medium := fmt.Sprintf("%s", row[6])
				adSetName := fmt.Sprintf("%s", row[7])
				budget := fmt.Sprintf("%s", row[8])
				status := fmt.Sprintf("%s", row[9])
				listing_id := fmt.Sprintf("%s", row[10])
				if status == "" {
					tableLeads = append(tableLeads, []string{dataIsComplete, name, email, phone, hub, project, medium, status, fmt.Sprintf("%d", positionInSpreadsheet), budget, adSetName, listing_id})
				}

			} else {
				if dataIsComplete == "EMAIL INCORRECTO" {
					continue
				}
				return tableLeads, nil
			}

		}
	}
	return tableLeads, nil
}

func (sc *ServiceClient) WriteSpreadSheet(spreadsheetId string, lead entity.Lead, col string, message string, sheet string) {
	writeRange := sheet + "!" + col + strconv.Itoa(lead.RowNumber)
	values := [][]interface{}{
		{message},
	}
	valueRange := &sheets.ValueRange{
		Values: values,
	}
	_, err := sc.Srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
}

func (sc *ServiceClient) ReadSpreadSheetChargeLeads(spreadsheetId string) ([][]string, error) {
	const dataWithoutPhone = "SE ENVIA SIN PHONE"
	const continueWithoutPhone = "PHONE INCORRECTO"
	const continueWithoutEmail = "EMAIL INCORRECTO"

	var tableLeads [][]string
	readRange := "leads!X2:AU"
	resp, err := sc.Srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		return tableLeads, nil
	} else {
		for number, row := range resp.Values {
			positionInSpreadsheet := number + 2
			dataIsComplete := fmt.Sprintf("%s", row[0])
			if dataIsComplete == "DATOS COMPLETOS" || dataIsComplete == dataWithoutPhone {
				name := fmt.Sprintf("%s", row[1])
				email := fmt.Sprintf("%s", row[2])
				phone := fmt.Sprintf("%s", row[3])
				hub := fmt.Sprintf("%s", row[4])
				location_of_interest_codes := fmt.Sprintf("%s", row[5])
				project := fmt.Sprintf("%s", row[6])
				projectId := fmt.Sprintf("%s", row[7])
				source := fmt.Sprintf("%s", row[8])
				medium := fmt.Sprintf("%s", row[9])
				campaignName := fmt.Sprintf("%s", row[10])
				campaignId := fmt.Sprintf("%s", row[11])
				adSetName := fmt.Sprintf("%s", row[12])
				adSetId := fmt.Sprintf("%s", row[13])
				adName := fmt.Sprintf("%s", row[14])
				adId := fmt.Sprintf("%s", row[15])
				budget := fmt.Sprintf("%s", row[16])
				monthly_payment_budget := fmt.Sprintf("%s", row[17])
				expected_purchase_time := fmt.Sprintf("%s", row[18])
				purchase_purpose := fmt.Sprintf("%s", row[19])
				preferred_property_condition := fmt.Sprintf("%s", row[20])

				status := fmt.Sprintf("%s", row[21])
				id := fmt.Sprintf("%s", row[22])
				if status == "" {
					tableLeads = append(tableLeads, []string{dataIsComplete, name, email, phone, hub, location_of_interest_codes, project, projectId, source, medium, campaignName, campaignId, adSetName, adSetId, adName, adId, budget, monthly_payment_budget, expected_purchase_time, purchase_purpose, preferred_property_condition, status, fmt.Sprintf("%d", positionInSpreadsheet), id})
				}

			} else {
				if dataIsComplete == "DATOS INCOMPLETOS" || dataIsComplete == continueWithoutEmail || dataIsComplete == continueWithoutPhone {
					continue
				}
				return tableLeads, nil
			}

		}
	}
	return tableLeads, nil
}
