package figma

import "time"

type versionResponse struct {
	Versions []Version `json:"versions"`
}

// Version describes a version of a file.
type Version struct {
	// Unique identifier for version
	ID string `json:"id"`

	// The label given to the version in the editor
	Label string `json:"label"`

	// The description of the version as entered in the editor
	Description string `json:"description"`

	// The user that created the version
	User User `json:"user"`

	// The time at which the version was created
	CreatedAt time.Time `json:"created_at"`
}
