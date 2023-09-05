package eventstream

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var errPageDeleteTest = errors.New("page delete test error")
var pageDeleteTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var pageDeleteTestSince = time.Now().UTC()
var pageDeleteTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	4656021: {
		Topic:     "eqiad.mediawiki.page-delete",
		PageTitle: "réduisit",
		RevID:     22058660,
	},
	4656283: {
		Topic:     "eqiad.mediawiki.page-delete",
		PageTitle: "récupéreraient",
		RevID:     22110162,
	},
}

const pageDeleteTestExecURL = "/page-delete-exec"
const pageDeleteTestSubURL = "/page-delete-sub"

func createPageDeleteServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("page-delete.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(pageDeleteTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			_, err = w.Write(stub)

			if err != nil {
				log.Panic(err)
			} else {
				f.Flush()
			}
		}
	})

	router.HandleFunc(pageDeleteTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			_, err = w.Write(stub)

			if err != nil {
				log.Panic(err)
			} else {
				f.Flush()
			}
		}
	})

	return router, nil
}

func testPageDeleteEvent(t *testing.T, evt *PageDelete) {
	expected := pageDeleteTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestPageDeleteExec(t *testing.T) {
	router, err := createPageDeleteServer(t, &pageDeleteTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestExecURL,
		}).
		Build()

	stream := client.PageDelete(context.Background(), pageDeleteTestSince, func(evt *PageDelete) error {
		testPageDeleteEvent(t, evt)
		return nil
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestPageDeleteSub(t *testing.T) {
	since := pageDeleteTestSince
	router, err := createPageDeleteServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.PageDelete(ctx, pageDeleteTestSince, func(evt *PageDelete) error {
		testPageDeleteEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}

		return nil
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), pageDeleteTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}

func TestPageDeleteExecError(t *testing.T) {
	since := pageDeleteTestSince
	router, err := createPageDeleteServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestSubURL,
		}).
		Build()

	stream := client.PageDelete(context.Background(), pageDeleteTestSince, func(evt *PageDelete) error {
		testPageDeleteEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errPageDeleteTest
	})

	assert.Equal(t, errPageDeleteTest, stream.Exec())
}

func TestPageDeleteSubError(t *testing.T) {
	since := pageDeleteTestSince
	router, err := createPageDeleteServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestSubURL,
		}).
		Build()

	stream := client.PageDelete(context.Background(), pageDeleteTestSince, func(evt *PageDelete) error {
		testPageDeleteEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errPageDeleteTest
	})

	for err := range stream.Sub() {
		assert.Equal(t, errPageDeleteTest, err)
		break
	}
}
