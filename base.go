package eventstream

type baseData struct {
	Schema    string    `json:"$schema"`
	Meta      Meta      `json:"meta"`
	Performer Performer `json:"performer"`
	RevID     int       `json:"rev_id"`
}

type baseSchema struct {
	ID   []Info
	Data struct {
		baseData
	}
}
