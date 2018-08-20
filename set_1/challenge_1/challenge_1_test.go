package challenge_1

import "testing"

func TestFirstChallenge(t *testing.T) {
	src := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	outputStr := string(hexToBase64(src))
	expectedStr := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if outputStr != expectedStr {
		t.Errorf("Error hex to base64")
	}

}
