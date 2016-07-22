package interfaces

//DbHandler -
type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

//Row -
type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

//DbRepo -
type DbRepo struct {
	DbHandler DbHandler
}
