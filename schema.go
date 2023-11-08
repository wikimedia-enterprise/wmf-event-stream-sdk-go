package eventstream

import (
	"time"
)

type schema interface {
	canaryEventGetter
	unmarshal(evt *Event) error
	timestamp() time.Time
}

type canaryEventGetter interface {
	isCanaryEvent() bool
}

func parseSchema(sch schema, msg *Event, store *storage) {
	if err := sch.unmarshal(msg); err != nil {
		store.setError(err)
	} else {
		store.setSince(sch.timestamp())
	}
}

func isCanaryEvent(sch schema) bool {
	return sch.isCanaryEvent()
}
