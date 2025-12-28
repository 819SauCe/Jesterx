package models

import (
	"database/sql"
	"time"
)

type Page struct {
	Id        string         `json:"id"`
	TenantId  string         `json:"tenant_id"`
	Name      string         `json:"name"`
	PageId    string         `json:"page_id"`
	Domain    sql.NullString `json:"domain"`
	ThemeId   sql.NullString `json:"theme_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
