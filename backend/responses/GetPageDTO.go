package responses

import "time"

type PageDTO struct {
	Id         string    `json:"id"`
	Tenant_id  string    `json:"tenant_id"`
	Name       string    `json:"name"`
	Page_id    string    `json:"page_id"`
	Domain     string    `json:"domain,omitempty"`
	Theme_id   string    `json:"theme_id,omitempty"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
