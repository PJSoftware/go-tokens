package gotokens_test

import (
	"testing"

	"github.com/pjsoftware/gotokens"
)

var fp []string = []string{"c:", "z:"}

func TestNoSearchPath(t *testing.T) {
	_, err := gotokens.ImportTokens("TEST/twitter-api.json")
	eec := gotokens.ENOSEARCHPATH
	if err == nil {
		t.Errorf("Expected '%s' error", eec)
	} else if gec := gotokens.ErrorCode(err); gec != eec {
		t.Errorf("Expected error code %s; got %s", eec, err)
	}
}

func TestImportTokens(t *testing.T) {
	gotokens.SetSearchPath(fp)
	_, err := gotokens.ImportTokens("TEST/twitter-api.json")
	if err != nil {
		t.Errorf("Error importing valid file: %s", err)
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
			gotEC := gotokens.ErrorCode(err)
			gotCtxt := gotokens.ErrorContext(err)
			if gotEC != tst.ec {
				t.Errorf("Expected error code %s; got %s: %s", tst.ec, gotEC, err)
			}
			if tst.ctxt != "" && gotCtxt != tst.ctxt {
				t.Errorf("Expected '%s' error context, got '%s': %s", tst.ctxt, gotokens.ErrorContext(err), err)
			}
		}
	}
}
