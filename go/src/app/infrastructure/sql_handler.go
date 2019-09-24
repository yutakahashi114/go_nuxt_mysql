package infrastructure

import (
	"app/interfaces/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_DATABASE")
	sock := os.Getenv("DB_SOCKET")
	dns := fmt.Sprintf("%s:%s@%s(%s)/%s", user, pass, sock, host, name)
	conn, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
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

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	row := new(SqlRow)
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return row, err
	}
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r *SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r *SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r *SqlRow) Close() error {
	return r.Rows.Close()
}
