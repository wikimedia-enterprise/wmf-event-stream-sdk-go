package eventstream

import (
	"encoding/json"
	"time"
)

// PageChange event scheme struct
type PageChange struct {
	ID   []Info
	Data struct {
		Schema    string `json:"$schema"`
		Meta      Meta   `json:"meta"`
		Performer struct {
			UserText           string    `json:"user_text"`
			UserGroups         []string  `json:"groups"`
			UserIsBot          bool      `json:"is_bot"`
			UserID             int       `json:"user_id"`
			UserRegistrationDt time.Time `json:"registration_dt"`
			UserEditCount      int       `json:"edit_count"`
		} `json:"performer"`
		Dt             time.Time `json:"dt"`
		ChangelogKind  string    `json:"changelog_kind"`
		PageChangeKind string    `json:"page_change_kind"`
		Page           struct {
			PageID         int64  `json:"page_id"`
			PageTitle      string `json:"page_title"`
			PageNamespace  int    `json:"namespace_id"`
			PageIsRedirect bool   `json:"is_redirect"`
		} `json:"page"`
		Revision struct {
			RevID        int64     `json:"rev_id"`
			RevDt        time.Time `json:"rev_dt"`
			Comment      string    `json:"comment"`
			ContentSlots struct {
				Main struct {
					ContentFormat string `json:"content_format"`
					ContentModel  string `json:"content_model"`
				} `json:"main"`
			} `json:"content_slots"`
			IsCommentVisible bool   `json:"is_comment_visible"`
			IsContentVisible bool   `json:"is_content_visible"`
			IsEditorVisible  bool   `json:"is_editor_visible"`
			IsMinorEdit      bool   `json:"is_minor_edit"`
			RevParentID      int64  `json:"rev_parent_id"`
			RevSha1          string `json:"rev_sha1"`
			RevSize          int    `json:"rev_size"`
		} `json:"revision"`
		PriorState struct {
			Page struct {
				PageTitle     string `json:"page_title"`
				PageNamespace int    `json:"namespace_id"`
			} `json:"page"`
			Revision struct {
				RevID int64 `json:"rev_id"`
			} `json:"revision"`
		} `json:"prior_state"`
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
