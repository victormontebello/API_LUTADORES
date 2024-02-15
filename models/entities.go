package models

type Lutador struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
	Peso int8    `json:"peso"`
	Titulos int8 `json:"titulos"`
	Ativo bool `json:"ativo"`
}