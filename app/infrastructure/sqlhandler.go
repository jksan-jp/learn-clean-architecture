package infrastructure

import (
	"database/sql"
	"learn-clean-architecture/interfaces/database"
)

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

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffectec() (int64, error) {
	return r.Result.RowsAffected()
}
