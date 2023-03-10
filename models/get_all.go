package models

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := db.QueryRow(`SELECT * FROM TODOS`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Subject, &todo.Content, &todo.Done)
		if err != nil {
			continue // ignora
		}

		todo = append(todos, todo)
	}

	return
}
