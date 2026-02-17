package models

import (
	"fmt"
)

func (result ChampionScores) Print() {
	for _, r := range result {
		fmt.Printf(
			"EventName: %s, AgeGroup: %s, GenderBracket: %s, FirstName: %s, LastName: %s, Total: %d\n",
			r.EventName, r.AgeGroup, r.GenderBracket, r.FirstName, r.LastName, r.Total,
		)
	}
}

type ChampionScores []*ChampionScoresTotal

type ChampionScoresTotal struct {
	EventName     string
	AgeGroup      string
	GenderBracket string
	FirstName     string
	LastName      string
	Total         int
}
