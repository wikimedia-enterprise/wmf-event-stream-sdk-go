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

var errPgPageChangeTest = errors.New("page change test error")
var pgPageChangeTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var pgPageChangeTestSince = time.Now().UTC()
var pgPageChangeTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	72231974: {
		Topic:     "eqiad.mediawiki.page-change",
		PageTitle: "User_talk:NR_01_RE",
		RevID:     1121302102,
	},
	9052925: {
		Topic:     "eqiad.mediawiki.page-change",
		PageTitle: "beaggiefa",
		RevID:     69852686,
	},
}

const pgPageChangeTestExecURL = "/page-change-exec"
const pgPageChangeTestSubURL = "/page-change-sub"

func createPageChangeServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("page-change.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(pgPageChangeTestExecURL, func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc(pgPageChangeTestSubURL, func(w http.ResponseWriter, r *http.Request) {
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

func testPgChangeEvent(t *testing.T, evt *PageChange) {
	expected := pgPageChangeTestResponse[evt.Data.PageID]

	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.Page.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.Revision.RevID)
}

func TestPgPageChangeExec(t *testing.T) {
	router, err := createPageChangeServer(t, &pgPageChangeTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageChangeURL: pgPageChangeTestExecURL,
		}).
		Build()

	stream := client.PageChange(context.Background(), pgCreateTestSince, func(evt *PageChange) error {
		testPgChangeEvent(t, evt)
		return nil
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestPgPageChangeSub(t *testing.T) {
	router, err := createPageChangeServer(t, &pgPageChangeTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageChangeURL: pgPageChangeTestSubURL,
		}).
		Build()

	ctx, cancel := context.WithCancel(context.Background())
	msgs := 0
	stream := client.PageChange(ctx, pgCreateTestSince, func(evt *PageChange) error {
		testPgChangeEvent(t, evt)
		msgs++

		if msgs > 3 {
			cancel()
		}

		return nil
	})

	assert.Equal(t, context.Canceled, stream.Sub())
}

func TestPgPageChangeSubError(t *testing.T) {
	router, err := createPageChangeServer(t, &pgPageChangeTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageChangeURL: pgPageChangeTestSubURL,
		}).
		Build()

	ctx, cancel := context.WithCancel(context.Background())
	msgs := 0
	stream := client.PageChange(ctx, pgCreateTestSince, func(evt *PageChange) error {
		testPgChangeEvent(t, evt)
		msgs++

		if msgs > 3 {
			cancel()
		}

		return errPgPageChangeTest
	})

	assert.Equal(t, errPgPageChangeTest, stream.Sub())
}

func TestPgPageChangeSubExecError(t *testing.T) {
	router, err := createPageChangeServer(t, &pgPageChangeTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageChangeURL: pgPageChangeTestSubURL,
		}).
		Build()

	ctx, cancel := context.WithCancel(context.Background())
	msgs := 0
	stream := client.PageChange(ctx, pgCreateTestSince, func(evt *PageChange) error {
		testPgChangeEvent(t, evt)
		msgs++

		if msgs > 3 {
			cancel()
		}

		return nil
	})

	assert.Equal(t, io.EOF, stream.Exec())
}
