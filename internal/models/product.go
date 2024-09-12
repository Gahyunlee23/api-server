package models

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`
}
