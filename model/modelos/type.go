package usuario

import "time"

type Usuario struct {
	ID          int64     `json:"Erros,omitempty"`
	Nome        string    `json:"Nome,omitempty"`
	UserName    string    `json:"UserName,omitempty"`
	Email       string    `json:"Email,omitempty"`
	Password    string    `json:"Password,omitempty"`
	DataCriacao time.Time `json:"DataCriacao,omitempty"`
}
