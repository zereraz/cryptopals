package challenge_2

import (
	"cryptopals/set_1/challenge_1"
	"testing"
)

func TestXor(t *testing.T) {
	fst := []byte("1c0111001f010100061a024b53535009181c")
	snd := []byte("686974207468652062756c6c277320657965")
	res := challenge_1.DecodeHex([]byte("746865206b696420646f6e277420706c6179"))
	if !bytesIsEqual(xorEqualLen(fst, snd), res) {
		t.Error("Error!")
	}
}

func bytesIsEqual(fst []byte, snd []byte) bool {
	if len(fst) != len(snd) {
		return false
	}
	for i := 0; i < len(fst); i++ {
		if fst[i] != snd[i] {
			return false
		}
	}
	return true
}
