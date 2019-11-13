package main

import "fmt"
import "sort"

func main() {
	currentLeaders := []int{}
	fmt.Printf("Current Leaders %v\n", currentLeaders)
	fmt.Printf("Leaders Mapped To ranking  %v\n", scoresToDenseRankings(currentLeaders))
	myScores := []int{1}
	fmt.Printf("My Scores %v\n", myScores)
	fmt.Printf("My rankings  %v\n", ladderRanking(currentLeaders, myScores))

}

//currenLeaders must be sorted in descenting order
//myScores should be shorted in ascenting order
//This is an O(n) solution
func ladderRanking(currentLeaders []int, myScores []int) []int {
	myRankings := make([]int, len(myScores))
	rankings := scoresToDenseRankings(currentLeaders)

	i := 0
	j := len(myScores) - 1

	if len(rankings) == 0 {
		for i, _ := range myRankings {
			myRankings[i] = 1
		}
		return myRankings
	}

	for j >= 0 && i < len(rankings) {
		myScore := myScores[j]

		if myScore >= currentLeaders[i] {
			myRankings[j] = rankings[i]
			j--
		} else {
			i++
		}
	}

	var worstRanking int
	if len(rankings) > 0 {
		worstRanking = rankings[len(rankings)-1] + 1
	}

	for j >= 0 {
		myRankings[j] = worstRanking
		j--
	}

	return myRankings
}

//This is an O(nlogn) solution
func ladderRankingNotOrdered(currentLeaders []int, myScores []int) []int {
	sort.Ints(currentLeaders)
	sort.Ints(myScores)
	currentLeaders = reverse(currentLeaders)

	myRankings := ladderRanking(currentLeaders, myScores)

	return myRankings
}

// Scores must be in descenting order
// e.g. scores = []int{110,110,80,60,60,30,25}
// scoresToDenseRankings(scores) returns []int{1,1,2,3,3,4,5}
func scoresToDenseRankings(scores []int) []int {
	if len(scores) == 0 {
		return []int{}
	}

	rankings := make([]int, len(scores))
	ranking := 1
	rankings[0] = ranking
	previousScore := scores[0]

	for i := 1; i < len(scores); i++ {
		currentScore := scores[i]

		if currentScore != previousScore {
			ranking += 1
		}
		rankings[i] = ranking
		previousScore = currentScore
	}

	return rankings
}

func reverse(a []int) []int {
	left := 0
	right := len(a) - 1
	for left < right {
		tmp := a[left]
		a[left] = a[right]
		a[right] = tmp
		left++
		right--
	}
	return a
}
