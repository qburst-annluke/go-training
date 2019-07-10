package connect

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  . "../config"
)

func ConnectDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    Host, Port, User, Password, DbName)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

	return db
}
