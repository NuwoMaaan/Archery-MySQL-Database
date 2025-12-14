package handlers

import (
	"database/sql"
)

type QueryHandler func(db *sql.DB, sql string) (QueryResult, error)

var Registry = map[string]QueryHandler{
	"GetArcherByID": func(db *sql.DB, sql string) (QueryResult, error) {
		return GetArcherByID(db, sql)
	},
	"GetEndCountTotal": func(db *sql.DB, sql string) (QueryResult, error) {
		return GetEndCountTotal(db, sql)
	},
	"GetEndCountOrderTotalByDate": func(db *sql.DB, sql string) (QueryResult, error) {
		return GetEndCountOrderTotalByDate(db, sql)
	},
	"GetEndCountOrderTotalByRoundNum": func(db *sql.DB, sql string) (QueryResult, error) {
		return GetEndCountOrderTotalByRoundNum(db, sql)
	},
	"GetRoundAndCategoryDetails": func(db *sql.DB, sql string) (QueryResult, error) {
		return GetRoundAndCategoryDetails(db, sql)
	},
}
