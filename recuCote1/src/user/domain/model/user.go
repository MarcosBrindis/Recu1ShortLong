package model

type User struct {
	ID   int    `json:"id"`
	Edad int    `json:"edad"`
	Name string `json:"name"`
	Sexo bool   `json:"sexo"`
}
