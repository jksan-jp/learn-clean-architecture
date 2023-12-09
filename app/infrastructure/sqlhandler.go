package infrastructure

import "database/sql"

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHander() *SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error())
	}
	sqlHander := new(SqlHandler)
	sqlHander.Conn = conn
	return sqlHander
}
