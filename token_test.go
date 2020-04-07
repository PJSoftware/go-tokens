package gotokens_test

import (
	"testing"

	"github.com/pjsoftware/gotokens"
)

func getToken(tName string) *gotokens.Token {
	tks, _ := gotokens.ImportTokens("C:/Tokens/TEST/twitter-api.json")
	tk, _ := tks.Select(tName)
	return tk
}

func TestCredentials(t *testing.T) {
	tk := getToken("Twitter_App_One")
	_, err := tk.Credential("NoSuchCred")
	testErrorCode(t, err, gotokens.EBADCREDENTIAL)

	got, err := tk.Credential("CONSUMER_KEY")
	exp := "abcde12345"
	if got != exp {
		t.Errorf("Expected '%s' but got '%s'", exp, got)
	}
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
