package eventstream

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const eventTestTitle = "event"
const eventTestTopic = "mediawiki.event.test"
const eventTestTimestamp = 1605631446001
const eventTestData = `{"title": "%s"}`
const eventTestID = `[{"topic":"%s","partition":0,"timestamp":%d},{"topic":"%s","partition":0,"offset":-1}]`

type eventTestPayload struct {
	Title string `json:"title"`
}

func TestEvent(t *testing.T) {
	payload := new(eventTestPayload)
	evt := new(Event)

	assert.Nil(t, evt.SetID("id: "+fmt.Sprintf(eventTestID, eventTestTopic, eventTestTimestamp, eventTestTopic)))
	assert.Equal(t, 2, len(evt.ID))
	assert.Equal(t, eventTestTimestamp, evt.ID[0].Timestamp)

	for _, id := range evt.ID {
		assert.Equal(t, eventTestTopic, id.Topic)
	}

	assert.Nil(t, evt.SetData("data: "+fmt.Sprintf(eventTestData, eventTestTitle)))
	assert.Nil(t, json.Unmarshal(evt.Data, payload))
	assert.Equal(t, eventTestTitle, payload.Title)
}
