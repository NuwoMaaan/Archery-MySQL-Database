package queries

import (
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/models"
	"github.com/NuwoMaaan/Archery-MySQL-Database/internal/repository"

	"database/sql"
)

type QueryHandler func(db *sql.DB) (models.QueryResult, error)

type QueryOption struct {
	Description string
	Handler     QueryHandler
}

var Options = []QueryOption{
	{
		Description: "Get archer by ID",
		Handler:     repository.GetArcherByID,
	},
	{
		Description: "Get end count by ID",
		Handler:     repository.GetEndCountTotal,
	},
	{
		Description: "Select Score By Round Order By Total",
		Handler:     repository.GetEndCountOrderTotalByTotal,
	},
	{
		Description: "Select Score By Round Order By Date",
		Handler:     repository.GetEndCountOrderTotalByDate,
	},
	{
		Description: "Select Score By Round Order By RoundNum",
		Handler:     repository.GetEndCountOrderTotalByRoundNum,
	},
	{
		Description: "search for round and catergory requirements by Round Name & Category Age",
		Handler:     repository.GetRoundAndCategoryDetails,
	},
	{
		Description: "Search archers score per all ends above certain value (50)",
		Handler:     repository.GetArcherScoreAboveValue,
	},
	{
		Description: "Search championship scores by particular championship event name",
		Handler:     repository.GetChampionScoresByEventID,
	},
	{
		Description: "Select Personal Best For Selected Round",
		Handler:     repository.GetPersonalBestByRound,
	},
	{
		Description: "Select Best Round For Particular Round Out Of All Archers",
		Handler:     repository.GetPersonalBestAllArchers,
	},
}
