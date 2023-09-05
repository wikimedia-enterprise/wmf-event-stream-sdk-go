package eventstream

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var subscribeTestSince = time.Now().UTC()

const subscribeTestTitle = "hello world"
const subscribeTestURL = "/subscribe"
const subscribeTestTime = 1605631446001
const subscribeTestTopic = "mediaiki.eventstream.test"
const subscribeTestMsgCount = 10

type subscribeTestData struct {
	Title string `json:"title"`
}

func createSubscribeServer(t *testing.T) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(subscribeTestURL, func(w http.ResponseWriter, r *http.Request) {
		f := w.(http.Flusher)

		assert.Equal(t, subscribeTestSince.Format(time.RFC3339), r.URL.Query().Get("since"))

		for i := 1; i <= subscribeTestMsgCount; i++ {
			msg := `event: message` + "\n"
			msg += fmt.Sprintf(`id: [{"topic":"%s","partition":0,"timestamp":%d},{"topic":"%s","partition":0,"offset":-1}]`+"\n", subscribeTestTopic, subscribeTestTime, subscribeTestTopic)
			msg += `data: ` + fmt.Sprintf(`{ "title": "%s" }`, subscribeTestTitle) + "\n"

			_, err := w.Write([]byte(msg))

			if err != nil {
				log.Panic(err)
			} else {
				f.Flush()
			}
		}
	})

	return router
}

func TestSubscribe(t *testing.T) {
	srv := httptest.NewServer(createSubscribeServer(t))
	defer srv.Close()

	ctx := context.Background()
	client := new(http.Client)
	msgs := 0

	err := subscribe(ctx, client, srv.URL+subscribeTestURL, subscribeTestSince, func(evt *Event) {
		assert.NotNil(t, evt)
		assert.Equal(t, len(evt.ID), 2)
		assert.Equal(t, evt.ID[0].Timestamp, subscribeTestTime)
		msgs++

		data := new(subscribeTestData)
		assert.Nil(t, json.Unmarshal(evt.Data, data))
		assert.Equal(t, subscribeTestTitle, data.Title)

		for _, id := range evt.ID {
			assert.Equal(t, subscribeTestTopic, id.Topic)
		}
	})

	assert.Equal(t, subscribeTestMsgCount, msgs)
	assert.Equal(t, err, io.EOF)
}
