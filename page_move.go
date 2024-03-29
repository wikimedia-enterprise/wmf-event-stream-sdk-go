package eventstream

import (
	"encoding/json"
	"time"
)

// PageMove event scheme struct
type PageMove struct {
	baseSchema
	Data struct {
		baseData
		PageID         int    `json:"page_id"`
		PageTitle      string `json:"page_title"`
		PageNamespace  int    `json:"page_namespace"`
		PageIsRedirect bool   `json:"page_is_redirect"`
		Database       string `json:"database"`
		RevID          int    `json:"rev_id"`
		PriorState     struct {
			PageTitle     string `json:"page_title"`
			PageNamespace int    `json:"page_namespace"`
			RevID         int    `json:"rev_id"`
		} `json:"prior_state"`
		Comment       string `json:"comment"`
		Parsedcomment string `json:"parsedcomment"`
	}
}

func (pm *PageMove) timestamp() time.Time {
	return pm.Data.Meta.Dt
}

func (pm *PageMove) unmarshal(evt *Event) error {
	pm.ID = evt.ID
	return json.Unmarshal(evt.Data, &pm.Data)
}
