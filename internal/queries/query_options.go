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
}
