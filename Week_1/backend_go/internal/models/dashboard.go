package models

import "time"

type LastStudySession struct {
	ID        uint      `json:"id"`
	GroupID   uint      `json:"group_id"`
	Group     Group     `json:"group"`
	CreatedAt time.Time `json:"created_at"`
}
