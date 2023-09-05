package eventstream

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var errStreamTest = errors.New("stream test error")
var streamTestSince = time.Now().UTC().Add(time.Hour * 2)

const streamTestBackoff = time.Millisecond * 1

func TestStream(t *testing.T) {
	thrownErrs := 0
	caughtErrs := 0
	storage := newStorage(streamTestSince, streamTestBackoff)

	stream := NewStream(storage, func(since time.Time) error {
		assert.Equal(t, streamTestSince, since)
		thrownErrs++

		if thrownErrs < keepAliveNumberOfErrors {
			return errStreamTest
		}

		return context.Canceled
	})

	for err := range stream.Sub() {
		caughtErrs++

		if thrownErrs >= keepAliveNumberOfErrors {
			assert.Equal(t, context.Canceled, err)
		} else {
			assert.Equal(t, errStreamTest, err)
		}
	}

	assert.Equal(t, keepAliveNumberOfErrors, thrownErrs)
	assert.Equal(t, keepAliveNumberOfErrors, caughtErrs)

	stream = NewStream(newStorage(streamTestSince, streamTestBackoff), func(since time.Time) error {
		assert.Equal(t, streamTestSince, since)
		return errStreamTest
	})

	assert.Equal(t, errStreamTest, stream.Exec())
}
