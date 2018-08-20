package challenge_3

import (
	"cryptopals/set_1/challenge_1"
	"unicode/utf8"
)

func BuildCorpus(text string) map[rune]float64 {
	corpus := make(map[rune]float64)

	for _, c := range text {
		corpus[c] += 1
	}
	total := utf8.RuneCountInString(text)
	for c := range corpus {
		corpus[c] /= float64(total)
	}
	return corpus
}

func ScoreText(text string, corpus map[rune]float64) float64 {
	var total float64
	for _, c := range text {
		total += corpus[c]
	}
	total /= float64(utf8.RuneCountInString(text))
	return total
}

func SingleByteXor(str []byte, c byte) []byte {
	for i, b := range str {
		str[i] = c ^ b
	}
	return str
}

func FindSingleXorKey(str []byte, corpus map[rune]float64) (float64, string) {
	var lastScore float64
	var res string

	for i := 0; i < 256; i++ {
		c := rune(i)
		out := SingleByteXor(challenge_1.DecodeHex(str), byte(c))
		currScore := ScoreText(string(out), corpus)
		if lastScore < currScore {
			lastScore = currScore
			res = string(out)
		}
	}
	return lastScore, res
}
