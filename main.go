package main

import (
	"fmt"
	"math"
	"math/rand"
)

var AllTeam map[string][]string = map[string][]string{
	"Team A": {"จั๊ม A",
		// "บิว A",
		"เบ้น A",
		"บุ๋ม A",
		"ต้น A",
		"ฟิสท์ A",
		"กีต้า A",
		"อุ้ม A"},
	"Team B": {"เก๋ง B",
		// "แอน B",
		"เหมา B",
		"แฟร์ B",
		"กัส B",
		"กร B",
		"อ้น B",
		"ปอ B",
		"มด B"},
	"Team C": {"เนี๊ยบ C",
		"บาส C",
		"แสน C",
		"ภูริ C",
		"ป่าน C",
		"เก่ง C",
		// "อาท C",
		"โก้ C",
		"ฝน C"},
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
		// sliceFraction := make([]int, 0, fraction)
		// for i := 0; i < fraction; i++ {
		// 	sliceFraction = append(sliceFraction, 1)
		// }

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
				// randIndex := rand.Intn(len(sliceFraction))
				// quotaPerTeam[teamName] += sliceFraction[randIndex]
				// sliceFraction = append(sliceFraction[:randIndex], sliceFraction[randIndex+1:]...)
				quotaPerTeam[teamName] += 1
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
		fmt.Printf("----->%s \n", key)
		for _, member := range value {
			fmt.Println(member)
		}
	}
}
