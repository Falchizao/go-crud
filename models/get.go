package models

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := db.QueryRow(`SELECT * FROM TODOS WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Subject, &todo.Content, &todo.Done)

	return
}
