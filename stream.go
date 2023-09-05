package eventstream

import (
	"time"
)

// NewStream create new stream instance
func NewStream(store *storage, handler func(since time.Time) error) *Stream {
	return &Stream{
		store,
		handler,
	}
}

// Stream this is for controlling steam execution
type Stream struct {
	store   *storage
	handler func(since time.Time) error
}

// Exec blocking execution stream
func (sm *Stream) Exec() error {
	go func() {
		if err := sm.handler(sm.store.getSince()); err != nil {
			sm.store.setError(err)
		}
	}()

	for err := range sm.store.getErrors() {
		return err
	}

	return nil
}

// Sub non blocking execution stream
func (sm *Stream) Sub() chan error {
	go keepAlive(sm.handler, sm.store)
	return sm.store.getErrors()
}
