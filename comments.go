package figma

import "time"

type commentResponse struct {
	Comments Comments `json:"comments"`
}

// Comments is a slice of Comment.
type Comments []Comment

// Comment is a comment or reply left by a user.
type Comment struct {
	//	Unique identifier for comment
	ID string

	// The absolute coordinates of where the comment is on the canvas
	ClientMeta Vector `json:"client_meta"`

	// The file in which the comment lives
	FileKey string `json:"file_key"`

	// If present, the id of the comment to which this is the reply
	ParentID string `json:"parent_id"`

	// The user who left the comment
	User User `json:"user"`

	// The time at which the comment was left
	CreatedAt time.Time `json:"created_at"`

	// If set, when the comment was resolved
	ResolvedAt time.Time `json:"resolved_at"`

	// Only set for top level comments. The number displayed with the comment in the UI
	OrderID int `json:"order_id,string"`
}
