package eventstream

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var storageTestSince = time.Now().UTC()

const storageTestBackoff = time.Millisecond * 1

func TestStorage(t *testing.T) {
	thrownErrs := 0
	caughtErrs := 0
	storage := newStorage(storageTestSince, storageTestBackoff)

	assert.NotNil(t, &storage.mu)
	assert.NotNil(t, storage.errs)
	assert.NotNil(t, storage.backoff)
	assert.NotNil(t, storage.errs)
	assert.Equal(t, storageTestSince, storage.since)
	assert.Equal(t, storageTestSince, storage.getSince())
	assert.Equal(t, storageTestBackoff, storage.backoff)
	assert.Equal(t, storageTestBackoff, storage.getBackoff())

	go func() {
		for err := range storage.getErrors() {
			assert.NotNil(t, err)
			caughtErrs++
		}
	}()

	thrownErrs++
	storage.setError(fmt.Errorf("test error"))
	storage.closeErrors()
	assert.Equal(t, thrownErrs, caughtErrs)

	since := time.Now().Add(2 * time.Hour)
	storage.setSince(since)
	assert.Equal(t, since, storage.since)
	assert.Equal(t, since, storage.getSince())
}
