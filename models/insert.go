package models

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.CreateConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO TODOS (subject, content, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Subject, todo.Content, todo.Done).Scan(&id)

	return
}
