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
}
