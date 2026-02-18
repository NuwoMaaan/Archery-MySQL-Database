package models

import "fmt"

type PersonalBestIndividual struct {
	PB_TOTAL int
}

func (r *PersonalBestIndividual) Print() {
	fmt.Println("RESULT:")
	fmt.Println("Personal Best:", r.PB_TOTAL)
}
