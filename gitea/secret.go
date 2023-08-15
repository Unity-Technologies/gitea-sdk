package gitea

import "time"

type Secret struct {
	// the secret's name
	Name string `json:"name"`
	// Date and Time of secret creation
	Created time.Time `json:"created_at"`
}
