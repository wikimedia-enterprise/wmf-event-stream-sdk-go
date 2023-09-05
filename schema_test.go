package eventstream

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var schemaTestSince = time.Now().UTC().Add(time.Hour * 5)
var schemaTestTimestamp = time.Now().UTC().Add(time.Hour * 1)

const schemaTestInfoTopic = "schema.test.topic"
const schemaTestInfoPartition = 2
const schemaTestInfoTimestamp = 1605631446001
const schemaTestInfoOffset = -1
const schemaTestBackoff = time.Millisecond * 1
const schemaTestTitle = "schema test title"
const schemaTestData = `{ "title": "%s" }`

type schemaTest struct {
	baseSchema
	Data struct {
		Title string
	}
}

func (s *schemaTest) timestamp() time.Time {
	return schemaTestTimestamp
}

func (s *schemaTest) unmarshal(evt *Event) error {
	s.ID = evt.ID
	return json.Unmarshal(evt.Data, &s.Data)
}

func TestSchema(t *testing.T) {
	storage := newStorage(schemaTestSince, schemaTestBackoff)
	schema := new(schemaTest)
	event := Event{
		[]Info{
			{
				schemaTestInfoTopic,
				schemaTestInfoPartition,
				schemaTestInfoTimestamp,
				schemaTestInfoOffset,
			},
		},
		[]byte(fmt.Sprintf(schemaTestData, schemaTestTitle)),
	}

	go func() {
		for err := range storage.getErrors() {
			assert.Error(t, err)
		}
	}()

	parseSchema(schema, &event, storage)

	assert.NotEqual(t, schemaTestSince, storage.getSince())
	assert.Equal(t, schemaTestTimestamp, storage.getSince())
	assert.Equal(t, schema.Data.Title, schemaTestTitle)
	assert.Equal(t, 1, len(schema.ID))

	for _, id := range schema.ID {
		assert.Equal(t, schemaTestInfoTopic, id.Topic)
		assert.Equal(t, schemaTestInfoPartition, id.Partition)
		assert.Equal(t, schemaTestInfoTimestamp, id.Timestamp)
		assert.Equal(t, schemaTestInfoOffset, id.Offset)
	}
}
