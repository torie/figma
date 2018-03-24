package figma

import (
	"encoding/json"
)

type imageResponse struct {
	Error  string `json:"err"`
	Images Images `json:"images"`
	Status int    `json:"status"`
}

// Images is a slice of rendered nodes.
type Images []Image

// UnmarshalJSON implements the Unmarshaler interface.
func (i *Images) UnmarshalJSON(b []byte) error {
	var imap map[string]string
	if err := json.Unmarshal(b, &imap); err != nil {
		return err
	}

	for nid, url := range imap {
		*i = append(*i, Image{nid, url})
	}
	return nil
}

// Image is a reference to a rendered node from Figma.
type Image struct {
	NodeID string
	URL    string
}
