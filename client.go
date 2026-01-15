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
		"",
		Metrics{enabled: false},
	}
}

// SetUserAgent sets a client with useragent.
func (c *Client) SetUserAgent(ua string) {
	c.userAgent = ua
}

// Client request client
type Client struct {
	url         string
	httpClient  *http.Client
	backoffTime time.Duration
	options     *Options
	userAgent   string
	metrics     Metrics
}

// PageCreate connect to page create stream
func (cl *Client) PageCreate(tx context.Context, since time.Time, handler func(evt *PageCreate) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, pageCreateURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageCreateURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(cl.options.PageCreateURL)
			cl.metrics.IncTotalEvents(cl.options.PageCreateURL)

			evt := new(PageCreate)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(tx context.Context, since time.Time, handler func(evt *PageDelete) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, pageDeleteURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageDeleteURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(ctx.Value(MetricLabelStream).(string))
			cl.metrics.IncTotalEvents(ctx.Value(MetricLabelStream).(string))

			evt := new(PageDelete)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(tx context.Context, since time.Time, handler func(evt *PageMove) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, pageMoveURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageMoveURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(ctx.Value(MetricLabelStream).(string))
			cl.metrics.IncTotalEvents(ctx.Value(MetricLabelStream).(string))

			evt := new(PageMove)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(tx context.Context, since time.Time, handler func(evt *RevisionCreate) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, revisionCreateURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionCreateURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(ctx.Value(MetricLabelStream).(string))
			cl.metrics.IncTotalEvents(ctx.Value(MetricLabelStream).(string))

			evt := new(RevisionCreate)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}

// RevisionVisibilityChange connect to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(tx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, revisionVisibilityChangeURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionVisibilityChangeURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(ctx.Value(MetricLabelStream).(string))
			cl.metrics.IncTotalEvents(ctx.Value(MetricLabelStream).(string))

			evt := new(RevisionVisibilityChange)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}

// PageChange connect to page change stream
func (cl *Client) PageChange(tx context.Context, since time.Time, handler func(evt *PageChange) error) *Stream {
	store := newStorage(since, cl.backoffTime)
	ctx := context.WithValue(tx, MetricLabelStream, pageChangeURL)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageChangeURL, store.getSince(), cl.userAgent, cl.metrics, func(msg *Event) {
			cl.metrics.IncAverageEvents(ctx.Value(MetricLabelStream).(string))
			cl.metrics.IncTotalEvents(ctx.Value(MetricLabelStream).(string))

			evt := new(PageChange)
			parseSchema(ctx, evt, msg, store, cl.metrics)

			if err := handler(evt); err != nil {
				cl.metrics.IncTotalErrors(ctx.Value(MetricLabelStream).(string), SeverityLabelValueMedium, "no", "no")
				store.setError(err)
			}
		})
	})
}
