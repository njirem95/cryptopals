package scoreboard

import "cryptopals/decoder"

type Score struct {
	Character byte
	Result string
	Score float64
}

var LetterFrequency = map[rune]float64{
	'a': 8.167,
	'b': 1.492,
	'c': 2.782,
	'd': 4.253,
	'e': 12.702,
	'f': 2.228,
	'g': 2.015,
	'h': 6.094,
	'i': 6.966,
	'j': 0.153,
	'k': 0.772,
	'l': 4.025,
	'm': 2.406,
	'n': 6.749,
	'o': 7.507,
	'p': 1.929,
	'q': 0.095,
	'r': 5.987,
	's': 6.327,
	't': 9.056,
	'u': 2.758,
	'v': 0.978,
	'w': 2.360,
	'x': 0.150,
	'y': 1.974,
	'z': 0.074,
	' ': 13.000,
}


func FindBestScore(input []byte) (Score, error) {
	var topScore Score
	for i := 0; i < 256; i++ {
		character := byte(i)

		var score float64
		result := decoder.Decode(string(input), character)


		for _, res := range result {
			frequency, ok := LetterFrequency[res]
			if ok {
				score += frequency
			}
		}

		if score > topScore.Score {
			topScore = Score { Character: character, Score: score, Result: result }
		}
	}

	return topScore, nil
}