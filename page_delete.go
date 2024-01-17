package eventstream

import (
	"encoding/json"
	"time"
)

// PageDelete event scheme struct
type PageDelete struct {
	baseSchema
	Data struct {
		baseData
		PageID         int    `json:"page_id"`
		PageTitle      string `json:"page_title"`
		PageNamespace  int    `json:"page_namespace"`
		PageIsRedirect bool   `json:"page_is_redirect"`
		Database       string `json:"database"`
		RevID          int    `json:"rev_id"`
		RevCount       int    `json:"rev_count"`
		Comment        string `json:"comment"`
		Parsedcomment  string `json:"parsedcomment"`
	}
}

func (pd *PageDelete) timestamp() time.Time {
	return pd.Data.Meta.Dt
}

func (pd *PageDelete) unmarshal(evt *Event) error {
	pd.ID = evt.ID
	return json.Unmarshal(evt.Data, &pd.Data)
}
