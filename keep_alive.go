package eventstream

import (
	"context"
	"errors"
	"time"
)

func keepAlive(handler func(since time.Time) error, store *storage) {
	for {
		err := handler(store.getSince())
		store.setError(err)

		if errors.Is(err, context.Canceled) {
			store.closeErrors()
			return
		}

		time.Sleep(store.getBackoff())
	}
}
