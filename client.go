package eventstream

import (
	"context"
	"net/http"
	"time"
)

const url = "https://stream.wikimedia.org"

const backoffTime = time.Second * 1

// All the available streams
const (
	pageCreateURL               = "/v2/stream/page-create"
	pageDeleteURL               = "/v2/stream/page-delete"
	pageMoveURL                 = "/v2/stream/page-move"
	revisionCreateURL           = "/v2/stream/revision-create"
	revisionVisibilityChangeURL = "/v2/stream/mediawiki.revision-visibility-change"
	pageChangeURL               = "/v2/stream/mediawiki.page_change.v1"
)

// NewClient creating new connection client
func NewClient() *Client {
	return &Client{
		url,
		new(http.Client),
		backoffTime,
		&Options{
			pageCreateURL,
			pageDeleteURL,
			pageMoveURL,
			revisionCreateURL,
			revisionVisibilityChangeURL,
			pageChangeURL,
		},
	}
}

// Client request client
type Client struct {
	url         string
	httpClient  *http.Client
	backoffTime time.Duration
	options     *Options
}

// PageCreate connect to page create stream
func (cl *Client) PageCreate(ctx context.Context, since time.Time, handler func(evt *PageCreate) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageCreateURL, store.getSince(), func(msg *Event) {
			evt := new(PageCreate)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *PageDelete) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageDeleteURL, store.getSince(), func(msg *Event) {
			evt := new(PageDelete)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *PageMove) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageMoveURL, store.getSince(), func(msg *Event) {
			evt := new(PageMove)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *RevisionCreate) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionCreateURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionCreate)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}

// RevisionVisibilityChange connect to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionVisibilityChangeURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionVisibilityChange)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}

// PageChange connect to page change stream
func (cl *Client) PageChange(ctx context.Context, since time.Time, handler func(evt *PageChange) error) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageChangeURL, store.getSince(), func(msg *Event) {
			evt := new(PageChange)
			parseSchema(evt, msg, store)

			if err := handler(evt); err != nil {
				store.setError(err)
			}
		})
	})
}
