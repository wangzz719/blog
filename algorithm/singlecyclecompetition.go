package main

import "fmt"

// 单循环比赛
func singleCycleCompetition(athletes []int) {
	athleteLen := len(athletes)
	fmt.Println("=================")
	if athleteLen <= 1 {
		fmt.Println("")
	}
	isOdd := athleteLen%2 == 1
	days := athleteLen - 1
	fakeLen := athleteLen
	if isOdd {
		athletes = append(athletes, 0)
		days = athleteLen
		fakeLen = athleteLen + 1
	}
	for i := 0; i < days; i++ {
		fmt.Print("day ", i+1, " ")
		for j := 0; j < fakeLen/2; j++ {
			//if newAthletes[j] != 0 && newAthletes[fakeLen-1-j] != 0 {
			fmt.Print(athletes[j], "-", athletes[fakeLen-1-j], " ")
			//}
		}
		fmt.Println()
		firstAthletes := athletes[0:1]
		newAthletes := athletes[fakeLen-1 : fakeLen]
		newAthletes = append(newAthletes, athletes[1:fakeLen-1]...)
		newAthletes = append(firstAthletes, newAthletes...)
		fmt.Println(newAthletes)
	}
}

// 单循环比赛 贝格尔算法
func singleCycleCompetitionBergerAlgorithm(athletes []int) {
	athleteLen := len(athletes)
	fmt.Println("=================")
	if athleteLen <= 1 {
		fmt.Println("")
	}
	isOdd := athleteLen%2 == 1
	days := athleteLen - 1
	fakeLen := athleteLen
	if isOdd {
		athletes = append(athletes, 0)
		days = athleteLen
		fakeLen = athleteLen + 1
	}
	moveStep := 0
	if athleteLen > 4 {
		moveStep = (athleteLen + athleteLen%2 - 4) / 2
	}
	for i := 0; i < days; i++ {
		fmt.Print("day ", i+1, " ")
		for j := 0; j < fakeLen/2; j++ {
			fmt.Print(athletes[j], "-", athletes[fakeLen-1-j], " ")
		}
		fmt.Println()
		athletes[0], athletes[fakeLen-1] = athletes[fakeLen-1], athletes[0]
		if i%2 == 0 {
			firstAthletes := athletes[0:1]
			newAthletes := athletes[fakeLen-moveStep-1 : fakeLen]
			newAthletes = append(newAthletes, athletes[1:fakeLen-moveStep-1]...)
			newAthletes = append(firstAthletes, newAthletes...)
			athletes = newAthletes
		} else {
			lastAthletes := athletes[fakeLen-1]
			newAthletes := athletes[fakeLen-moveStep-2 : fakeLen-1]
			newAthletes = append(newAthletes, athletes[0:fakeLen-moveStep-2]...)
			newAthletes = append(newAthletes, lastAthletes)
			athletes = newAthletes
		}
		fmt.Println(athletes)
	}
}

func main() {
	singleCycleCompetition([]int{1, 2, 3, 4, 5, 6, 7})
	singleCycleCompetitionBergerAlgorithm([]int{1, 2, 3, 4, 5, 6, 7, 8})
}
