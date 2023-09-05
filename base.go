package eventstream

type baseData struct {
	Schema         string    `json:"$schema"`
	Database       string    `json:"database"`
	Meta           Meta      `json:"meta"`
	Performer      Performer `json:"performer"`
	PageID         int       `json:"page_id"`
	PageTitle      string    `json:"page_title"`
	PageNamespace  int       `json:"page_namespace"`
	PageIsRedirect bool      `json:"page_is_redirect"`
	RevID          int       `json:"rev_id"`
}

type baseSchema struct {
	ID   []Info
	Data struct {
		baseData
	}
}
