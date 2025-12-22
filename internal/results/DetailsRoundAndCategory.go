package results

import (
	"database/sql"
)

type DetailsRoundAndCategory struct {
	RoundID       int
	RoundName     string
	AgeGroup      string
	GenderBracket string
	BowType       string
	Distance10m   sql.NullString
	Distance20m   sql.NullString
	Distance30m   sql.NullString
	Distance40m   sql.NullString
	Distance50m   sql.NullString
	Distance60m   sql.NullString
	Distance70m   sql.NullString
	Distance90m   sql.NullString
}
