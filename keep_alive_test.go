package eventstream

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var errKeepAliveTest = errors.New("keep alive error")

const keepAliveTestBackoffTime = time.Millisecond * 1
const keepAliveNumberOfErrors = 5

func TestKeepAlive(t *testing.T) {
	storageSince := time.Now().UTC()
	thrownErrs := 0
	caughtErrs := 0
	storage := newStorage(storageSince, keepAliveTestBackoffTime)

	assert.NotNil(t, storage)

	handler := func(since time.Time) error {
		assert.Equal(t, storageSince, since)
		thrownErrs++

		if thrownErrs < keepAliveNumberOfErrors {
			return errKeepAliveTest
		}

		return context.Canceled
	}

	go keepAlive(handler, storage)

	for err := range storage.getErrors() {
		caughtErrs++

		if thrownErrs >= keepAliveNumberOfErrors {
			assert.Equal(t, context.Canceled, err)
		} else {
			assert.Equal(t, errKeepAliveTest, err)
			storageSince = storageSince.Add(time.Hour * 1)
			storage.setSince(storageSince)
		}
	}

	assert.Equal(t, keepAliveNumberOfErrors, thrownErrs)
	assert.Equal(t, keepAliveNumberOfErrors, caughtErrs)
}
