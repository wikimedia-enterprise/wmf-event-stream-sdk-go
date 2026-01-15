package eventstream

import (
	"context"
	"time"
)

type schema interface {
	unmarshal(evt *Event) error
	timestamp() time.Time
}

func parseSchema(ctx context.Context, sch schema, msg *Event, store *storage, metrics Metrics) {
	if err := sch.unmarshal(msg); err != nil {
		metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), "high", "yes", "yes")
		store.setError(err)
	} else {
		store.setSince(sch.timestamp())
	}
}
