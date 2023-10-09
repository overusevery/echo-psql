package entity

import "time"

type Todo struct {
	ID        int64
	Content   string
	Status    bool
	UpdatedAt time.Time
	CreatedAt time.Time
}
