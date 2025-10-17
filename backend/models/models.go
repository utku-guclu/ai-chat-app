package models

import "time"

// User represents a chat user.
// In our simplified app, users are created on first message send.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

// Message represents a single chat message, either from a human or the AI.
type Message struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	IsAI      bool      `json:"is_ai"` // True if message is from the AI bot
	CreatedAt time.Time `json:"created_at"`
	// The following field is for joining/display purposes only
	// and is not stored directly in the messages table.
	Username string `json:"username"`
}
