package request_lead

import (
	"github.com/gookit/goutil/structs"
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
	extraFields := structs.ToMap(event.Context)
	context := &analytics.Context{
		Extra: extraFields,
	}
	return sc.Client.Enqueue(analytics.Track{
		Event:       event.Event,
		AnonymousId: event.AnonymousID,
		Properties:  structs.ToMap(event.Properties),
		Context:     context,
	})
	/* return sc.Client.Enqueue(analytics.Track{
		Event:       "Lead Creation Requested",
		AnonymousId: "121456864645631455646",
		Properties:  analytics.NewProperties().Set("Full_Name", "John Doe"),
		//Context:     &context,
	}) */

}
