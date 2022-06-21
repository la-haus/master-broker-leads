package validator

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
)

func ValidatePhone(phone string, country string, config *configs.Config) string {
	client := &http.Client{}
	method := "GET"
	url := config.CustomerDataPlatform.Host + "/v1/customer-shield/validator?phone=" + phone + "&country-code=" + country
	payload := bytes.NewBuffer([]byte{})
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return phone
	}
	req.Header.Add("x-api-key", config.CustomerDataPlatform.ApiKey)
	resp, err := client.Do(req)
	if err != nil {
		return phone
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return phone
	}
	data := new(entity.PhoneResponse)
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return phone
	}
	return data.Phone
}
