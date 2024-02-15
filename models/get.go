package models

import "lutadores-modulo/db"

func Get(id int) (lutador Lutador, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow("SELECT id, nome, peso, titulos, ativo FROM lutadores WHERE id = $1", id)

	err = row.Scan(&lutador.ID, &lutador.Nome, &lutador.Peso, &lutador.Titulos, &lutador.Ativo)
	
	return 
}
