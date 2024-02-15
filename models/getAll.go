package models

import "lutadores-modulo/db"

func GetAll() (lutadores []Lutador, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT id, nome, peso, titulos, ativo FROM lutadores")

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var lutador Lutador
		err = rows.Scan(&lutador.ID, &lutador.Nome, &lutador.Peso, &lutador.Titulos, &lutador.Ativo)

		if err != nil {
			continue
		}

		lutadores = append(lutadores, lutador)
	}

	return
}