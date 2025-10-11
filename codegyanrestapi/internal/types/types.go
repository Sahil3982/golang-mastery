package types

type Student struct {
	Id   int `json:"id"`
	Name string `validate:"required"`
	Email string	`validate:"required"`
	Age int `validate:"required"`
}