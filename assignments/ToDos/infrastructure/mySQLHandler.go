package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //load mysql driver
	"github.com/vinod4006/TeamPolyglot-GO/assignments/ToDos/interfaces"
)

//MySqlHandler - connection
type MySqlHandler struct {
	Conn *sql.DB
}

func (handler *MySqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MySqlHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(MySqlRow)
	}
	row := new(MySqlRow)
	row.Rows = rows
	return row
}

type MySqlRow struct {
	Rows *sql.Rows
}

func (r MySqlRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r MySqlRow) Next() bool {
	return r.Rows.Next()
}

func NewMysqlHandler(connstring string) *MySqlHandler {
	conn, _ := sql.Open("mysql", connstring)
	sqlHandler := new(MySqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
