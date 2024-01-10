package easy

import "log"

func isWinner(player1, player2 []int) int {
	var score1, score2 = 0, 0
	for ind, score := range player1 {
		if (ind > 0 && player1[ind-1] == 10) || (ind > 1 && player1[ind-2] == 10) {
			score *= 2
		}
		score1 += score
	}

	for ind, score := range player2 {
		if (ind > 0 && player2[ind-1] == 10) || (ind > 1 && player2[ind-2] == 10) {
			score *= 2
		}
		score2 += score
	}
	log.Println("player 1:", score1, " player2", score2)
	if score1 > score2 {
		return 1
	} else if score1 == score2 {
		return 0
	}
	return 2
}
