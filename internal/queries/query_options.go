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
		Description: "Score By Round Order By Total",
		SQL:         "SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum FROM end INNER JOIN round on round.RoundID = end.RoundID INNER JOIN archer on archer.ArcherID = end.ArcherID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20' GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date ORDER BY RoundTotalSum DESC LIMIT 10",
		Handler:     handlers.GetEndCountOrderTotalByTotal,
	},
	{
		Description: "Select Score By Round Order By Date",
		SQL:         "SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum FROM end INNER JOIN round on round.RoundID = end.RoundID INNER JOIN archer on archer.ArcherID = end.ArcherID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20' GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date ORDER BY Date DESC LIMIT 10",
		Handler:     handlers.GetEndCountOrderTotalByDate,
	},
	{
		Description: "Select Score By Round Order By RoundNum",
		SQL:         "SELECT RoundNum, round.RoundName, archer.ArcherID, archer.FirstName, Date, archer.LastName, SUM(total) AS RoundTotalSum FROM end INNER JOIN round on round.RoundID = end.RoundID INNER JOIN archer on archer.ArcherID = end.ArcherID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, round.RoundName, Date ORDER BY RoundNum DESC LIMIT 10",
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
		Description: "Search championship scores by particular championship event name",
		SQL:         "SELECT archerevent.EventName, catergory.AgeGroup, catergory.GenderBracket, archer.FirstName, archer.LastName, SUM(TOTAL) as Total FROM end INNER JOIN archer ON archer.ArcherID = end.ArcherID INNER JOIN archerevent on archerevent.EventID = end.EventID INNER JOIN catergory on catergory.CatergoryID = end.CatergoryID WHERE archerevent.EventName = 'December Championship' AND Approved = 1 GROUP BY end.ArcherID, catergory.AgeGroup, catergory.GenderBracket;",
		Handler:     handlers.GetChampionScoresByEventID,
	},
	{
		Description: "Select Personal Best For Selected Round",
		SQL:         "SELECT MAX(round_total_sum) AS PB_total FROM (SELECT RoundNum, SUM(total) AS round_total_sum FROM end INNER JOIN round on round.RoundID = end.RoundID WHERE round.RoundName = 'Hobart' AND Approved = 1 AND ArcherID = 1  AND round.DateAdded IS NOT NULL and round.DateChange IS NULL GROUP BY RoundNum) AS round_sums",
		Handler:     handlers.GetPersonalBestByRound,
	},
	{
		Description: "Select Best Round For Particular Round Out Of All Archers",
		SQL:         "SELECT archer.ArcherID, archer.FirstName, archer.LastName, MAX(round_total_sum) AS PB_total FROM (SELECT RoundNum, archer.ArcherID, archer.FirstName, archer.LastName, SUM(total) AS round_total_sum FROM end INNER JOIN round on round.RoundID = end.RoundID INNER JOIN archer on archer.ArcherID = end.ArcherID WHERE round.RoundName = 'Hobart' AND Approved = 1 AND round.DateAdded IS NOT NULL AND round.DateChange IS NULL GROUP BY RoundNum, archer.ArcherID, archer.FirstName, archer.LastName) AS round_sums INNER JOIN  archer on archer.ArcherID = round_sums.ArcherID GROUP BY archer.ArcherID",
		Handler:     handlers.GetPersonalBestAllArchers,
	},
}
