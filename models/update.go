package models

import "crud/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec("UPDATE todo SET subject=$1, content=$2, done=$3 WHERE id=$4", todo.Subject, todo.Content, todo.Done, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
