package request_lead

import (
	"encoding/json"
	"fmt"

	"github.com/la-haus/master-brokers-charge-leads/configs"
	"github.com/la-haus/master-brokers-charge-leads/domain/entity"
	"github.com/segmentio/analytics-go/v3"
)

type SegmentClient struct {
	Client analytics.Client
}

func NewSegmentClient(config *configs.Config) *SegmentClient {

	return &SegmentClient{
		Client: analytics.New(config.SegmentWritekey),
	}
}

func (sc *SegmentClient) SendTrackLead(event entity.LeadCreationRequested) error {
	var properties map[string]interface{}
	var extraFields map[string]interface{}
	conData, err := json.Marshal(event.Context)
	if err != nil {
		fmt.Println("Error marshal context: ", err, event.Properties.Email)
		return err
	}
	err = json.Unmarshal(conData, &extraFields)
	if err != nil {
		fmt.Println("Error unmarshal context: ", err, event.Properties.Email)
		return err
	}
	context := &analytics.Context{
		Extra: extraFields,
	}
	data, err := json.Marshal(event.Properties) // Convert to a json string
	if err != nil {
		fmt.Println("Error marshal properties: ", err, event.Properties.Email)
		return err
	}

	err = json.Unmarshal(data, &properties) // Convert to a map
	if err != nil {
		fmt.Println("Error unmarshal properties: ", err, event.Properties.Email)
		return err
	}

	return sc.Client.Enqueue(analytics.Track{
		Event:       event.Event,
		AnonymousId: event.AnonymousID,
		Properties:  properties,
		Context:     context,
	})

}
