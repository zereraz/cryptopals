package challenge_1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func DecodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst
}

func hexToBase64(src []byte) []byte {
	dst := DecodeHex(src)
	base64Dst := make([]byte, base64.StdEncoding.EncodedLen(len(dst)))
	base64.StdEncoding.Encode(base64Dst, dst)
	return base64Dst
}
