package eventstream

import (
	"time"
)

type schema interface {
	unmarshal(evt *Event) error
	timestamp() time.Time
	hasCanaryEvent() bool
}

func parseSchema(sch schema, msg *Event, store *storage) {
	if err := sch.unmarshal(msg); err != nil {
		store.setError(err)
	} else {
		store.setSince(sch.timestamp())
	}
}

func hasCanaryEvent(sch schema) bool {
	return sch.hasCanaryEvent()
}
