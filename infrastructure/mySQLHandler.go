package infrastructure

import (
	"book-shelve/TeamPolyglot-GO/interfaces"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //load mysql driver
)

//MySQLHandler - connection
type MySQLHandler struct {
	Conn *sql.DB
}

//Execute -
func (handler *MySQLHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

//Query -
func (handler *MySQLHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(MySQLRow)
	}
	row := new(MySQLRow)
	row.Rows = rows
	return row
}

//MySQLRow -
type MySQLRow struct {
	Rows *sql.Rows
}

//Scan -
func (r MySQLRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

//Next -
func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}

//NewMySQLHandler -
func NewMySQLHandler(connstring string) *MySQLHandler {
	conn, _ := sql.Open("mysql", connstring)
	sqlHandler := new(MySQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
