package eventstream

import (
	"encoding/json"
	"time"
)

// RevisionCreate event scheme struct
type PageChange struct {
	baseSchema
	Data struct {
		baseData
		Dt             time.Time `json:"dt"`
		ChangelogKind  string    `json:"changelog_kind"`
		PageChangeKind string    `json:"page_change_kind"`
		Page           struct {
			PageID    int    `json:"page_id"`
			PageTitle string `json:"page_title"`
		} `json:"page"`
		Revision struct {
			RevID        int       `json:"rev_id"`
			RevDt        time.Time `json:"rev_dt"`
			Comment      string    `json:"comment"`
			ContentSlots struct {
				ContentFormat string `json:"content_format"`
				ContentModel  string `json:"content_model"`
			} `json:"content_slots"`
			IsCommentVisible bool   `json:"is_comment_visible"`
			IsContentVisible bool   `json:"is_content_visible"`
			IsEditorVisible  bool   `json:"is_editor_visible"`
			IsMinorEdit      bool   `json:"is_minor_edit"`
			RevParentID      int    `json:"rev_parent_id"`
			RevSha1          string `json:"rev_sha1"`
			RevSize          int    `json:"rev_size"`
		} `json:"revision"`
		Database string `json:"wiki_id"`
	}
}

func (rc *PageChange) timestamp() time.Time {
	return rc.Data.Meta.Dt
}

func (rc *PageChange) unmarshal(evt *Event) error {
	rc.ID = evt.ID
	return json.Unmarshal(evt.Data, &rc.Data)
}
