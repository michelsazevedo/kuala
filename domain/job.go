package domain

import "time"

//Job ...
type Job struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CompanyId   int64     `json:"company_id" pg:"pk_id" validate:"required"`
	Tags        []string  `json:"tags" pg:",array"`
	Featured    bool      `json:"featured"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at" pg:"default:now()"`
	ExpiresAt   time.Time `json:"expires_at"`
	DeletedAt   time.Time `json:"-" pg:", soft_delete"`
}
