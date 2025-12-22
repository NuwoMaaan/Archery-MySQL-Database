package results

type QueryResult interface {
	Print(sql string)
}
