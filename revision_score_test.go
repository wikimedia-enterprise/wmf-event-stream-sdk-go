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

var errRevScoreTest = errors.New("revision score test error")
var revScoreTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var revScoreTestSince = time.Now().UTC()
var revScoreTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	66132507: {
		Topic:     "eqiad.mediawiki.revision-score",
		PageTitle: "Q66533108",
		RevID:     1316923273,
	},
	13731303: {
		Topic:     "eqiad.mediawiki.revision-score",
		PageTitle: "Utilisateur:Denvis1/NCAA-Squelette_équipe",
		RevID:     177205614,
	},
}

const revScoreTestExecURL = "/revision-score-exec"
const revScoreTestSubURL = "/revision-score-sub"

func createRevScoreServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("revision-score.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(revScoreTestExecURL, func(w http.ResponseWriter, r *http.Request) {
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

	router.HandleFunc(revScoreTestSubURL, func(w http.ResponseWriter, r *http.Request) {
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

func testRevScoreEvent(t *testing.T, evt *RevisionScore) {
	expected := revScoreTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestRevScoreExec(t *testing.T) {
	router, err := createRevScoreServer(t, &revScoreTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revScoreTestExecURL,
		}).
		Build()

	stream := client.RevisionScore(context.Background(), revScoreTestSince, func(evt *RevisionScore) error {
		testRevScoreEvent(t, evt)
		return nil
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestRevScoreSub(t *testing.T) {
	since := revScoreTestSince
	router, err := createRevScoreServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revScoreTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.RevisionScore(ctx, revScoreTestSince, func(evt *RevisionScore) error {
		testRevScoreEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}

		return nil
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), revScoreTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}

func TestRevisionScoreExecError(t *testing.T) {
	since := revScoreTestSince
	router, err := createRevScoreServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revScoreTestExecURL,
		}).
		Build()

	stream := client.RevisionScore(context.Background(), revScoreTestSince, func(evt *RevisionScore) error {
		testRevScoreEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errRevScoreTest
	})

	assert.Equal(t, errRevScoreTest, stream.Exec())
}

func TestRevisionScoreSubError(t *testing.T) {
	since := revScoreTestSince
	router, err := createRevScoreServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revScoreTestSubURL,
		}).
		Build()

	stream := client.RevisionScore(context.Background(), revScoreTestSince, func(evt *RevisionScore) error {
		testRevScoreEvent(t, evt)
		since = evt.Data.Meta.Dt
		return errRevScoreTest
	})

	for err := range stream.Sub() {
		assert.Equal(t, errRevScoreTest, err)
		break
	}
}
