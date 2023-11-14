package eventstream

import (
	"encoding/json"
	"time"
)

// RevisionCreate event scheme struct
type RevisionCreate struct {
	baseSchema
	Data struct {
		baseData
		RevTimestamp      time.Time `json:"rev_timestamp"`
		RevSha1           string    `json:"rev_sha1"`
		RevMinorEdit      bool      `json:"rev_minor_edit"`
		RevLen            int       `json:"rev_len"`
		RevContentModel   string    `json:"rev_content_model"`
		RevContentFormat  string    `json:"rev_content_format"`
		Comment           string    `json:"comment"`
		ChronologyID      string    `json:"chronology_id"`
		Parsedcomment     string    `json:"parsedcomment"`
		RevParentID       int       `json:"rev_parent_id"`
		RevContentChanged bool      `json:"rev_content_changed"`
		RevIsRevert       bool      `json:"rev_is_revert"`
	}
}

func (rc *RevisionCreate) timestamp() time.Time {
	return rc.Data.Meta.Dt
}

func (rc *RevisionCreate) unmarshal(evt *Event) error {
	rc.ID = evt.ID
	return json.Unmarshal(evt.Data, &rc.Data)
}
