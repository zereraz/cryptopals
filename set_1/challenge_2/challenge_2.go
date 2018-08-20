package challenge_2

import "cryptopals/set_1/challenge_1"

func xorEqualLen(fst []byte, snd []byte) []byte {
	fst = challenge_1.DecodeHex(fst)
	snd = challenge_1.DecodeHex(snd)
	if len(fst) != len(snd) {
		panic("different length for both buffers")
	}
	output := make([]byte, len(fst))
	for i := 0; i < len(fst); i++ {
		output[i] = fst[i] ^ snd[i]
	}
	return output
}
