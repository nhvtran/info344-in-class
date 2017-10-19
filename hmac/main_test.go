package main

import "testing"

func TestSign(t *testing.T) {
	//TODO: write unit test cases for sign()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader

	// echo "a value to sign" | hmac sign somesecretkey
	// >> returns the digital signature in the terminal
	cases := []struct {
		input          string
		signingKey     string
		expectedOutput string
	}{
		{
			input:          "hello",
			signingKey:     "somesecretkey",
			expectedOutput: "8nnOkxiw7ywvBiTTpAisbXIfEat9ru3HNBrnQmB7aMU=",
		},
	}

	for _, c := range cases {

	}
}

func TestVerify(t *testing.T) {
	//TODO: write until test cases for verify()
	//use strings.NewReader() to get an io.Reader
	//interface over a simple string
	//https://golang.org/pkg/strings/#NewReader
}
