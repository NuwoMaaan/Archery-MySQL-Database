package queries

import (
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/handlers"
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/results"

	"database/sql"
)

type QueryHandler func(db *sql.DB, sql string) (results.QueryResult, error)

type QueryOption struct {
	Description string
	SQL         string
	Handler     QueryHandler
}

var Options = []QueryOption{
	{
		Description: "Get archer by ID",
		SQL:         "SELECT FirstName, LastName, Gender FROM archer WHERE ArcherId = 1",
		Handler:     handlers.GetArcherByID,
	},
	{
		Description: "Get end count by ID",
		SQL:         "SELECT ArcherID, COUNT(ArcherID) FROM end GROUP BY ArcherID",
		Handler:     handlers.GetEndCountTotal,
	},
	{
		Description: "Search archers scores per end between set dates ORDER BY TOTAL",
		SQL:         "SELECT endID, RoundNum, FirstName, LastName, round.RoundName, Date, Distance, TargetType, TOTAL FROM archer INNER JOIN end ON archer.ArcherID = end.ArcherID INNER JOIN round on round.RoundID = end.RoundID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20' ORDER BY TOTAL DESC LIMIT 15",
		Handler:     handlers.GetEndCountOrderTotalByDate,
	},
	{
		Description: "Search Archers score per end between set dates ORDER BY RoundNum",
		SQL:         "SELECT endID, RoundNum, FirstName, LastName, round.RoundName, Date, Distance, TargetType, TOTAL FROM archer INNER JOIN end ON archer.ArcherID = end.ArcherID INNER JOIN round on round.RoundID = end.RoundID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20'ORDER BY RoundNum DESC LIMIT 15",
		Handler:     handlers.GetEndCountOrderTotalByRoundNum,
	},
	{
		Description: "search for round and catergory requirements by Round Name & Category Age",
		SQL:         "SELECT round.RoundID, round.RoundName, catergory.AgeGroup, catergory.GenderBracket, catergory.BowType, 10m,20m,30m,40m,50m,60m,70m,90m FROM catergory INNER JOIN round on round.RoundID = catergory.RoundID WHERE round.RoundName = 'Perth' AND catergory.AgeGroup = 'Under 21'",
		Handler:     handlers.GetRoundAndCategoryDetails,
	},
	{
		Description: "Search archers score per all ends above certain value (50)",
		SQL:         "SELECT end.ArcherID, COUNT(TOTAL) FROM end LEFT JOIN archer ON archer.ArcherID = end.ArcherID WHERE TOTAL > 50 AND Approved = 1 GROUP BY end.ArcherID",
		Handler:     handlers.GetArcherScoreAboveValue,
	},
	{
		Description: "Search championship scores using inner joins of Archer, End and Category tables",
		SQL:         "SELECT archerevent.EventName, catergory.AgeGroup, catergory.GenderBracket, archer.FirstName, archer.LastName, SUM(TOTAL) as Total FROM end INNER JOIN archer ON archer.ArcherID = end.ArcherID INNER JOIN archerevent on archerevent.EventID = end.EventID INNER JOIN catergory on catergory.CatergoryID = end.CatergoryID WHERE end.EventID = 603 AND Approved = 1 GROUP BY end.ArcherID, catergory.AgeGroup, catergory.GenderBracket;",
		Handler:     handlers.GetChampionScoresByEventID,
	},
}
