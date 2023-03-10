package db

import (
	"database/sql"
	"fmt"

	_ "github.com/Falchizao/go-crud/crud/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.getDBConfig()

	sc := fmt.Sprintf("host=%s port=% user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
