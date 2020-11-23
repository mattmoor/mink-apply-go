package divide

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/mattmoor/mink-apply-go/types"
)

func handle(req types.Payload) types.Payload {
	return types.Payload{
		A: req.A / req.B,
		B: req.A % req.B,
	}
}

func Receiver(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, error) {
	// Parse the payload that we receive.
	req := types.Payload{}
	if err := event.DataAs(&req); err != nil {
		return nil, cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
	}

	// Manipulate the payload
	resp := handle(req)

	// Respond with a new cloudevent with the mutated payload.
	r := cloudevents.NewEvent(cloudevents.VersionV1)
	r.SetType("dev.mink.apply.samples.divide")
	r.SetSource("https://github.com/mattmoor/mink-apply-go/divide")
	if err := r.SetData(event.DataContentType(), resp); err != nil {
		return nil, cloudevents.NewHTTPResult(500, "failed to set response data: %s", err)
	}
	return &r, nil
}
