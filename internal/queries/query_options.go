package queries

type QueryOption struct {
	Description string
	SQL         string
	HandlerKey  string
}

var Options = []QueryOption{
	{
		Description: "Get archer by ID",
		SQL:         "SELECT FirstName, LastName, Gender FROM archer WHERE ArcherId = 1",
		HandlerKey:  "GetArcherByID",
	},
	{
		Description: "Get end count by ID",
		SQL:         "SELECT ArcherID, COUNT(ArcherID) FROM end GROUP BY ArcherID",
		HandlerKey:  "GetEndCountTotal",
	},
	{
		Description: "Search archers scores per end between set dates ORDER BY TOTAL",
		SQL:         "SELECT endID, RoundNum, FirstName, LastName, round.RoundName, Date, Distance, TargetType, TOTAL FROM archer INNER JOIN end ON archer.ArcherID = end.ArcherID INNER JOIN round on round.RoundID = end.RoundID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20' ORDER BY TOTAL DESC LIMIT 15",
		HandlerKey:  "GetEndCountOrderTotalByDate",
	},
	{
		Description: "Search Archers score per end between set dates ORDER BY RoundNum",
		SQL:         "SELECT endID, RoundNum, FirstName, LastName, round.RoundName, Date, Distance, TargetType, TOTAL FROM archer INNER JOIN end ON archer.ArcherID = end.ArcherID INNER JOIN round on round.RoundID = end.RoundID WHERE archer.ArcherID = 1 AND Approved = 1 AND round.RoundName = 'Hobart' AND Date between '2023-04-01' and '2023-04-20'ORDER BY RoundNum DESC LIMIT 15",
		HandlerKey:  "GetEndCountOrderTotalByRoundNum",
	},
	{
		Description: "search for round and catergory requirements by Round Name & Category Age",
		SQL:         "SELECT round.RoundID, round.RoundName, catergory.AgeGroup, catergory.GenderBracket, catergory.BowType, 10m,20m,30m,40m,50m,60m,70m,90m FROM catergory INNER JOIN round on round.RoundID = catergory.RoundID WHERE round.RoundName = 'Perth' AND catergory.AgeGroup = 'Under 21'",
		HandlerKey:  "GetRoundAndCategoryDetails",
	},
}
