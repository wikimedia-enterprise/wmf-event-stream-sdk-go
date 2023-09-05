package eventstream

import (
	"time"
)

type schema interface {
	unmarshal(evt *Event) error
	timestamp() time.Time
}

func parseSchema(sch schema, msg *Event, store *storage) {
	if err := sch.unmarshal(msg); err != nil {
		store.setError(err)
	} else {
		store.setSince(sch.timestamp())
	}
}
