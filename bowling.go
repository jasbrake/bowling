package bowling

import (
	"log"
	"strconv"
)

// ScoreGame calculates the final score of a bowling game.
// The game is passed in as a string of frames separated by a '-' (without
// quotation marks). Strikes are represented with an 'X', spares are represented
// with a '/', and an open frame is two numerical digits.
// We assume that the game string passed in is a valid game.
func ScoreGame(game string) int {
	total := 0
	// the number of throws so far. 20 is the last non-bonus throw
	throw := 0
	// bonus tells us whether we apply a bonus from a strike or spare or not
	bonus := 0
	// store the previous number to calculate the spare value easily
	prevNumber := 0

	for i := 0; i < len(game); i++ {
		// ignore separation chars
		if game[i] == '-' {
			continue
		}

		// the bonus has to be added later so it doesn't affect this round
		additionalBonus := 0
		num := 0
		switch game[i] {
		case 'X':
			num = 10
			additionalBonus = 2
			throw += 2
			break
		case '/':
			num = 10 - prevNumber
			additionalBonus = 1
			throw++
			break
		default:
			n, err := strconv.Atoi(string(game[i]))
			if err != nil {
				log.Fatalf("%b is an invalid bowling throw", game[i])
			}
			num = n
			prevNumber = n
			throw++
		}

		if throw <= 20 {
			total += num
		}

		if bonus > 0 {
			var calcBonus int
			if bonus > 2 {
				calcBonus = 2 * num
				bonus -= 2
			} else {
				calcBonus = num
				bonus--
			}
			total += calcBonus
		}
		if throw <= 20 {
			bonus += additionalBonus
			additionalBonus = 0
		}
	}
	return total
}
