package models

import (
	"lutadores-modulo/db"
)

func Delete(id int) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "DELETE FROM lutadores WHERE id = $1"

	resultado, err := conn.Exec(sql, id)

	return resultado.RowsAffected()
}