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
		Dt            time.Time `json:"dt"`
		ChangelogKind string    `json:"changelog_kind"`
		Page          struct {
			PageID    int    `json:"page_id"`
			PageTitle string `json:"page_title"`
		}
		Revision struct {
			RevID int       `json:"rev_id"`
			RevDt time.Time `json:"rev_dt"`
		}
		WikiID string `json:"wiki_id"`
	}
}

func (rc *PageChange) timestamp() time.Time {
	return rc.Data.Meta.Dt
}

func (rc *PageChange) unmarshal(evt *Event) error {
	rc.ID = evt.ID
	return json.Unmarshal(evt.Data, &rc.Data)
}
