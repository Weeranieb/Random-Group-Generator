package main

import (
	"fmt"
	"math"
	"math/rand"
)

var AllTeam map[string][]string = map[string][]string{
	"Team A": {
		"Name1 A",
		"Name2 A",
		"Name3 A",
		"Name4 A",
		"Name5 A",
		"Name6 A",
	},
	"Team B": {
		"Name1 B",
		"Name2 B",
		"Name3 B",
		"Name4 B",
		"Name5 B",
	},
	"Team C": {
		"Name1 C",
		"Name2 C",
		"Name3 C",
		"Name4 C",
		"Name5 C",
		"Name6 C",
		"Name7 C",
	},
}

var NewTeam map[string][]string = map[string][]string{
	"Team Alpha":   {},
	"Team Bravo":   {},
	"Team Charlie": {},
}

func main() {
	// Call the function

	var mapCountMember map[string]int = map[string]int{}

	for _, member := range AllTeam {
		isCountNotEqual := false
		maxLenInTeam := 0
		totalNewMember := 0
		for key, value := range NewTeam {
			mapCountMember[key] = len(value)
			totalNewMember += len(value)
		}

		if totalNewMember%len(NewTeam) != 0 {
			isCountNotEqual = true
			maxLenInTeam = int(math.Ceil(float64(totalNewMember) / float64(len(NewTeam))))
		}

		tempMember := member
		fraction := len(member) % len(AllTeam) // หารเอาเศษ
		quotient := len(member) / len(AllTeam) // ผลหาร
		sliceFraction := make([]int, 0, fraction)
		for i := 0; i < fraction; i++ {
			sliceFraction = append(sliceFraction, 1)
		}

		quotaPerTeam := map[string]int{}
		// make everyone equal as possible
		if fraction > 0 && isCountNotEqual {
			for tempTeam, countMember := range mapCountMember {
				if countMember < maxLenInTeam {
					quotaPerTeam[tempTeam] = 1
					fraction--
				}

				if fraction == 0 {
					break
				}
			}
		}

		for teamName := range NewTeam {
			quotaPerTeam[teamName] += quotient
			// random fraction ว่าจะเอาไปใส่ทีมไหน
			if fraction > 0 {
				randIndex := rand.Intn(len(sliceFraction))
				quotaPerTeam[teamName] += sliceFraction[randIndex]
				sliceFraction = append(sliceFraction[:randIndex], sliceFraction[randIndex+1:]...)
				fraction--
			}
		}

		for teamName, quota := range quotaPerTeam {
			for i := 0; i < quota; i++ {
				// random tempTeamMember
				randIndex := rand.Intn(len(tempMember))
				selectMember := tempMember[randIndex]

				// add member to new team
				NewTeam[teamName] = append(NewTeam[teamName], selectMember)

				// remove member from tempTeamMember
				tempMember = append(tempMember[:randIndex], tempMember[randIndex+1:]...)
			}
		}
	}

	for key, value := range NewTeam {
		fmt.Printf("-----> %s: %d person(s)\n", key, len(value))
		for _, member := range value {
			fmt.Println(member)
		}
	}
}
