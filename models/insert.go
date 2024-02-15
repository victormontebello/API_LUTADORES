package models

import (
	"lutadores-modulo/db"
)

func Insert(lutador Lutador) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := "INSERT INTO lutadores (nome, peso, titulos, ativo) VALUES ($1, $2, $3, $4) RETURNING id"

	err = conn.QueryRow(sql, lutador.Nome, lutador.Peso, lutador.Titulos, lutador.Ativo).Scan(&id)

	return
}