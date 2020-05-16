package bowling

import (
	"log"
	"strconv"
)

// NormalThrowCount is the number of throws in a game not counting bonus throws
const NormalThrowCount = 20

// StrikeScore is the number of points a strike is worth
const StrikeScore = 10

// ScoreGame calculates the final score of a bowling game.
// The game is passed in as a string of frames each separated by a '-' (without
// quotation marks). Strikes are represented with an 'X', spares are represented
// with a '/', and an open frame is two digits representing the number of pins
// knocked down by the two throws in that frame.
//
// We assume that the game string passed in is a valid game.
func ScoreGame(game string) int {
	total := 0
	// the number of throws so far
	throw := 0
	// bonus tells us whether we apply a bonus from a previous strike or spare
	// and will always fall in the range [0,3]
	bonus := 0
	// used to calculate a spare's point value easily
	prevNumber := 0

	for i := 0; i < len(game); i++ {
		// ignore separation chars
		if game[i] == '-' {
			continue
		}

		currentThrowBonus := 0
		num := 0
		switch game[i] {
		case 'X':
			num = StrikeScore
			currentThrowBonus = 2
			throw += 2
			break
		case '/':
			num = StrikeScore - prevNumber
			currentThrowBonus = 1
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

		if bonus > 0 {
			// there's a 2x bonus if the previous two throws were strikes
			if bonus >= 3 {
				total += 2 * num
				bonus -= 2
			} else {
				total += num
				bonus--
			}
		}

		// bonus throws don't get counted as regular throws
		if throw <= NormalThrowCount {
			total += num
			bonus += currentThrowBonus
		}
	}
	return total
}
