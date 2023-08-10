package models

type UpdatePatchRequest struct {
	Id     string                 `json:"id"`
	Fields map[string]interface{} `json:"fields"`
}
