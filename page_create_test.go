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

var errPgCreateTest = errors.New("page create test error")
var pgCreateTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var pgCreateTestSince = time.Now().UTC()
var pgCreateTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	72231974: {
		Topic:     "eqiad.mediawiki.page-create",
		PageTitle: "User_talk:NR_01_RE",
		RevID:     1121302102,
	},
	9052925: {
		Topic:     "eqiad.mediawiki.page-create",
		PageTitle: "beaggiefa",
		RevID:     69852686,
	},
}

const pgCreateTestExecURL = "/page-create-exec"
const pgCreateTestSubURL = "/page-create-sub"

func createPgCreateServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("page-create.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(pgCreateTestExecURL, func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc(pgCreateTestSubURL, func(w http.ResponseWriter, r *http.Request) {
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

func testPgCreateEvent(t *testing.T, evt *PageCreate) {
	expected := pgCreateTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestPgCreateExec(t *testing.T) {
	router, err := createPgCreateServer(t, &pgCreateTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageCreateURL: pgCreateTestExecURL,
		}).
		Build()

	stream := client.PageCreate(context.Background(), pgCreateTestSince, func(evt *PageCreate) error {
		testPgCreateEvent(t, evt)
		return nil
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestPageCreateSub(t *testing.T) {
	since := pgCreateTestSince
	router, err := createPgCreateServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageCreateURL: pgCreateTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.PageCreate(ctx, pgCreateTestSince, func(evt *PageCreate) error {
		testPgCreateEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}

		return nil
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), pgCreateTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}

func TestPageCreateExecError(t *testing.T) {
	since := pgCreateTestSince
	router, err := createPgCreateServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageCreateURL: pgCreateTestSubURL,
		}).
		Build()

	stream := client.PageCreate(context.Background(), pgCreateTestSince, func(evt *PageCreate) error {
		testPgCreateEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errPgCreateTest
	})

	assert.Equal(t, errPgCreateTest, stream.Exec())
}

func TestPageCreateSubError(t *testing.T) {
	since := pgCreateTestSince
	router, err := createPgCreateServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageCreateURL: pgCreateTestSubURL,
		}).
		Build()

	stream := client.PageCreate(context.Background(), pgCreateTestSince, func(evt *PageCreate) error {
		testPgCreateEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errPgCreateTest
	})

	for err := range stream.Sub() {
		assert.Equal(t, errPgCreateTest, err)
		break
	}
}
