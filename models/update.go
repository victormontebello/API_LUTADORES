package models

import "lutadores-modulo/db"

func Update(id int, lutador Lutador) (int64 , error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := "UPDATE lutadores SET nome = $1, peso = $2, titulos = $3, ativo = $4 WHERE id = $5"

	res, err := conn.Exec(sql, lutador.Nome, lutador.Peso, lutador.Titulos, lutador.Ativo, id)

	if  err != nil {
		return 0, err		
	}

	return res.RowsAffected()
}