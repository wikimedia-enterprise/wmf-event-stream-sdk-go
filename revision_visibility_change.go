package eventstream

import (
	"encoding/json"
	"time"
)

// RevisionVisibilityChange event scheme struct
type RevisionVisibilityChange struct {
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
		RevParentID      int       `json:"rev_parent_id"`
		Visibility       struct {
			Text    bool `json:"text"`
			User    bool `json:"user"`
			Comment bool `json:"comment"`
		} `json:"visibility"`
		PriorState struct {
			Visibility struct {
				Text    bool `json:"text"`
				User    bool `json:"user"`
				Comment bool `json:"comment"`
			} `json:"visibility"`
		} `json:"prior_state"`
	}
}

func (rvc *RevisionVisibilityChange) timestamp() time.Time {
	return rvc.Data.Meta.Dt
}

func (rvc *RevisionVisibilityChange) unmarshal(evt *Event) error {
	rvc.ID = evt.ID
	return json.Unmarshal(evt.Data, &rvc.Data)
}
