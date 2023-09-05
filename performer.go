package eventstream

import "time"

// Performer user that is responsible for event
type Performer struct {
	UserText           string    `json:"user_text"`
	UserGroups         []string  `json:"user_groups"`
	UserIsBot          bool      `json:"user_is_bot"`
	UserID             int       `json:"user_id"`
	UserRegistrationDt time.Time `json:"user_registration_dt"`
	UserEditCount      int       `json:"user_edit_count"`
}
