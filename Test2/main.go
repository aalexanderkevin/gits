/*
	GITS sedang bermain permainan arcade, dan dalam setiap permainan GITS ingin naik
	ke peringkat tertinggi dan juga ingin mengetahui setiap peringkat di setiap permainan.
	Dalam permainan ini menggunakan skema Dense Ranking dan memiliki aturan sebagai
	berikut:
		○ Peringkat pertama dapat diraih oleh pemain yang memiliki skor tertinggi.
		○ Pemain yang memiliki skor yang sama memiliki peringkat yang sama.
	Contoh :
		○ Empat pemain memiliki skor tertinggi sebagai berikut 100, 80, 80, dan 70, maka
			masing-masing pemain itu memiliki rangking 1,2,2 dan 3. Jika GITS memiliki skor
			60, 70, 100 setelah pertandingan maka rangking yang didapatkan adalah 4, 3
			dan 1.
	Buatlah program untuk menghitung rangking pemain ?
		sample input :
			7
			100 100 50 40 40 20 10
			4
			5 25 50 120
		sample output :
			6 4 2 1
*/

package main

import (
	"fmt"
	"sort"
)

func denseRank(scores []int, playerScores []int) []int {
	rankMap := make(map[int]int)
	rank := 1

	for _, score := range scores {
		if _, ok := rankMap[score]; !ok {
			rankMap[score] = rank
			rank++
		}
	}

	var result []int
	for i, playerScore := range playerScores {
		for iScore, score := range scores {
			if playerScore > score {
				result = append(result, iScore+1)
				break
			} else if _, exist := rankMap[playerScore]; exist {
				result = append(result, rankMap[playerScore])
				break
			}
		}
		if len(result) < i+1 {
			result = append(result, len(scores)+1)
		}
	}
	return result
}

func main() {
	var n, m int
	fmt.Scan(&n)

	scores := make([]int, 0)
	scoreMap := make(map[int]bool)

	for i := 0; i < n; i++ {
		var score int
		fmt.Scan(&score)
		if _, ok := scoreMap[score]; !ok {
			scoreMap[score] = true
			scores = append(scores, score)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	fmt.Scan(&m)
	playerScores := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&playerScores[i])
	}

	ranks := denseRank(scores, playerScores)

	for _, rank := range ranks {
		fmt.Printf("%d ", rank)
	}
}
