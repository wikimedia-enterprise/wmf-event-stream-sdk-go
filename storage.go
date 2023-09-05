package eventstream

import (
	"sync"
	"time"
)

func newStorage(since time.Time, backoff time.Duration) *storage {
	return &storage{
		sync.Mutex{},
		since,
		backoff,
		make(chan error),
	}
}

type storage struct {
	mu      sync.Mutex
	since   time.Time
	backoff time.Duration
	errs    chan error
}

func (st *storage) getErrors() chan error {
	return st.errs
}

func (st *storage) setError(err error) {
	st.errs <- err
}

func (st *storage) closeErrors() {
	close(st.errs)
}

func (st *storage) getSince() time.Time {
	return st.since
}

func (st *storage) setSince(since time.Time) {
	st.mu.Lock()
	st.since = since
	st.mu.Unlock()
}

func (st *storage) getBackoff() time.Duration {
	if st.backoff == 0 {
		st.backoff = time.Second * 1
	}

	return st.backoff
}
