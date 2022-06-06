package google_func

import (
	"fmt"
	"log"
	"strconv"

	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
	"google.golang.org/api/sheets/v4"
)

func ReadSpreadSheet(srv *sheets.Service, spreadsheetId string) ([][]string, error) {
	//spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
	var tableLeads [][]string
	readRange := "leads!K2:T"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
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
				status := ""
				if len(row) == 10 {
					status = fmt.Sprintf("%s", row[9])
				}
				if status == "" {
					tableLeads = append(tableLeads, []string{dataIsComplete, name, email, phone, hub, project, medium, status, fmt.Sprintf("%d", positionInSpreadsheet), budget, adSetName})
				}

			} else {
				return tableLeads, nil
			}
		}
	}
	return tableLeads, nil
}

func WriteSpreadSheet(srv *sheets.Service, spreadsheetId string, lead entity.Lead, isSend bool) {
	if isSend {
		writeRange := "leads!T" + strconv.Itoa(lead.RowNumber)
		values := [][]interface{}{
			{"Enviado"},
		}
		valueRange := &sheets.ValueRange{
			Values: values,
		}
		_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
	} else {
		writeRange := "leads!R" + strconv.Itoa(lead.RowNumber)
		values := [][]interface{}{
			{"Error"},
		}
		valueRange := &sheets.ValueRange{
			Values: values,
		}
		_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve data from sheet: %v", err)
		}
	}
}
