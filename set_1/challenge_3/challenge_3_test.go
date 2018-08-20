package challenge_3

import (
	"io/ioutil"
	"testing"
)

func TestSingleByte(t *testing.T) {
	str := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	corpusBytes, err := ioutil.ReadFile("./beowolf.txt")
	if err != nil {
		t.Fatal(err)
	}
	corpus := BuildCorpus(string(corpusBytes))
	_, res := FindSingleXorKey(str, corpus)
	t.Logf(res)
}
