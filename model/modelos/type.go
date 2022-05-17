package usuario

import "time"

type Usuario struct {
	ID          int64
	Nome        string
	Email       string
	UserName    string
	DataCriaçao time.Time
	Password    string
}
