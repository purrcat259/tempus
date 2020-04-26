package models

import "time"

// Entry - represents a single time entry
type Entry struct {
	ID      string
	Start   time.Time
	End     *time.Time
	Message *string
}
