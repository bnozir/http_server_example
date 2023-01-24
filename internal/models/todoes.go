package models

import "time"

type Todoes struct {
	ID             int
	Name           string
	StartTime      *time.Time
	ExpirationTime *time.Time
}
