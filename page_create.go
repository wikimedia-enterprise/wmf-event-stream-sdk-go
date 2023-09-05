package eventstream

import (
	"encoding/json"
	"time"
)

// PageCreate event schema struct
type PageCreate struct {
	baseSchema
	Data struct {
		baseData
		RevTimestamp     time.Time `json:"rev_timestamp"`
		RevSha1          string    `json:"rev_sha1"`
		RevMinorEdit     bool      `json:"rev_minor_edit"`
		RevLen           int       `json:"rev_len"`
		RevContentModel  string    `json:"rev_content_model"`
		RevContentFormat string    `json:"rev_content_format"`
		Comment          string    `json:"comment"`
		Parsedcomment    string    `json:"parsedcomment"`
	}
}

func (pc *PageCreate) timestamp() time.Time {
	return pc.Data.Meta.Dt
}

func (pc *PageCreate) unmarshal(evt *Event) error {
	pc.ID = evt.ID
	return json.Unmarshal(evt.Data, &pc.Data)
}
