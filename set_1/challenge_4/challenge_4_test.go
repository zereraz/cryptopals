package challenge_4

import (
	"cryptopals/set_1/challenge_3"
	"io/ioutil"
	"strings"
	"testing"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSingleByte(t *testing.T) {
	data, err := ioutil.ReadFile("./challenge-data.txt")
	corpusBytes, err := ioutil.ReadFile("../challenge_3/beowolf.txt")

	if err != nil {
		t.Fatal(err)
	}

	corpus := challenge_3.BuildCorpus(string(corpusBytes))
	var lastScore float64
	var res string

	for _, line := range strings.Split(string(data), "\n") {

		score, out := challenge_3.FindSingleXorKey([]byte(line), corpus)
		if score > lastScore {
			lastScore = score
			res = string(out)
		}
	}
	t.Log(res)
}
