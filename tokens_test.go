package gotokens_test

import (
	"testing"

	"github.com/pjsoftware/gotokens"
)

var fp []string = []string{"c:", "z:"}

func TestNoSearchPath(t *testing.T) {
	_, err := gotokens.ImportTokens("TEST/twitter-api.json")
	exp := gotokens.ENOSEARCHPATH
	testErrorCode(t, err, exp)
}

func TestImportTokens(t *testing.T) {
	gotokens.SetSearchPath(fp)
	_, err := gotokens.ImportTokens("TEST/twitter-api.json")
	if err != nil {
		t.Errorf("Error importing valid file: %s", err)
	}
}

func TestImportDirectly(t *testing.T) {
	path := "C:/Tokens/TEST/twitter-api.json"
	tks, err := gotokens.ImportTokens(path)
	if err != nil {
		t.Errorf("Error importing valid file directly: %s", err)
	}
	if tks.File() != path {
		t.Errorf("Tokens filename was '%s'; expected '%s'", tks.File(), path)
	}
}

func TestTokensExceptions(t *testing.T) {
	gotokens.SetSearchPath(fp)
	type tType struct {
		fn   string
		desc string
		ctxt string
		ec   string
	}
	tTable := []tType{
		{"dup-token", "Duplicate token names", "dup-token", gotokens.EMALFORMEDJSON},
		{"missing-tokens", "Missing 'tokens' key", "no-token", gotokens.EMALFORMEDJSON},
		{"missing-name", "Missing 'name' key", "no-name", gotokens.EMALFORMEDJSON},
		{"missing-cred", "Missing 'credentials' key", "no-cred", gotokens.EMALFORMEDJSON},
		{"nosuch-file", "Missing json file", "NoContext", gotokens.EFILENOTFOUND},
		{"not-json", "Badly formatted JSON", "json.Unmarshal", gotokens.EINTERNAL},
	}
	for _, tst := range tTable {
		tks, err := gotokens.ImportTokens("TEST/" + tst.fn + ".json")
		if err == nil {
			t.Errorf("%s (%s) error expected", tst.desc, tks.File())
		} else {
			testErrorCode(t, err, tst.ec)
			testErrorContext(t, err, tst.ctxt)
		}
	}
}

func TestChooseToken(t *testing.T) {
	tokens, err := gotokens.ImportTokens("C:/Tokens/TEST/twitter-api.json")
	if err != nil {
		t.Errorf("Error importing valid file directly: %s", err)
		return
	}

	_, err = tokens.Select("Twitter_App_One")
	if err != nil {
		t.Errorf("Error selecting existing token: %s", err)
	}

	_, err = tokens.Select("NoSuchToken")
	if err != nil {
		testErrorCode(t, err, gotokens.EBADTOKEN)
	}
}
