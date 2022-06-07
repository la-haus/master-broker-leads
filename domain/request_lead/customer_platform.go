package request_lead

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/la-haus/master-brokers-charge-leads/configs"
)

func SendEvent(event []byte, config *configs.Config) error {
	client := &http.Client{}
	method := "POST"
	//url := "https://customer-platform.staging.lahaus.com/v1/events"
	url := config.CustomerDataPlatform.Host
	payload := bytes.NewBuffer(event)
	req, err := http.NewRequest(method, url, payload)
	fmt.Println(req, err)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", config.CustomerDataPlatform.ApiKey)

	resp, err := client.Do(req)
	fmt.Println(req, err)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return err
	}
	return nil
}
