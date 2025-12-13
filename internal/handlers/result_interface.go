package handlers

type QueryResult interface {
	Print(sql string)
}
