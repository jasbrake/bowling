package bowling

import "testing"

type gameTest struct {
	game  string
	score int
}

func TestScoreGame(t *testing.T) {
	tests := []gameTest{
		{
			game:  "X-X-X-X-X-X-X-X-X-XXX",
			score: 300,
		},
		{
			game:  "45-54-36-27-09-63-81-18-90-72",
			score: 90,
		},
		{
			game:  "5/-5/-5/-5/-5/-5/-5/-5/-5/-5/-5",
			score: 150,
		},
		{
			game:  "63-2/-X-5/-71-22-X-X-90-00",
			score: 135,
		},
	}

	for _, test := range tests {
		score := ScoreGame(test.game)
		if test.score != score {
			t.Errorf("Game %s failed: expected score %d but calculated score %d", test.game, test.score, score)
		}
	}
}
