package usuario

import "time"

type Usuario struct {
	ID          int64     `json:"Id,omitempty"`
	Nome        string    `json:"Nome,omitempty"`
	UserName    string    `json:"UserName,omitempty"`
	Email       string    `json:"Email,omitempty"`
	Password    string    `json:"Password,omitempty"`
	NewPassword string    `json:"NewPassword,omitempty"`
	DataCriacao time.Time `json:"DataCriacao,omitempty"`
}
