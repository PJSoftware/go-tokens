package gotokens

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Tokens represents n token sets found in a token.json file
type Tokens struct {
	file   string
	tokens map[string]*Token
}

// Token represents an individual token set extracted from Tokens
type Token struct {
	cred map[string]string
}

var tokenSearchPath []string = nil

// SetSearchPath takes a slice of path strings to be searched for the Tokens folder
func SetSearchPath(fp []string) {
	tokenSearchPath = fp
}

// ImportTokens creates a new token data source from the specified file
func ImportTokens(fn string) (*Tokens, error) {
	tf, err := findJSONFile(fn)
	if err != nil {
		return nil, err
	}

	tks := new(Tokens)
	tks.file = tf
	tks.tokens = make(map[string]*Token)

	err = tks.read()
	if err != nil {
		return nil, err
	}
	return tks, nil
}

// File returns the name of the file the tokens were read from
func (tks *Tokens) File() string {
	return tks.file
}

func (tks *Tokens) read() error {
	jsonFile, _ := os.Open(tks.file)
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	return tks.parse(bytes)
}

func (tks *Tokens) parse(bytes []byte) error {
	data, err := convert(bytes)
	if err != nil {
		return err
	}

	for _, token := range data.([]interface{}) {
		tm := token.(map[string]interface{})

		var tn, ci interface{}
		var ok bool
		if tn, ok = tm["name"]; !ok {
			return &Error{
				Code:    EMALFORMEDJSON,
				Message: fmt.Sprintf("Malformed file: 'name' missing"),
			}
		}
		if ci, ok = tm["credentials"]; !ok {
			return &Error{
				Code:    EMALFORMEDJSON,
				Message: fmt.Sprintf("Malformed file: 'credentials' missing"),
			}
		}

		cs := ci.([]interface{})
		csi := cs[0]
		cm := csi.(map[string]interface{}) // silently swallows duplicate credential entries
		tk := new(Token)
		tns := tn.(string)
		if _, ok = tks.tokens[tns]; ok {
			return &Error{
				Code:    EMALFORMEDJSON,
				Message: fmt.Sprintf("Malformed file: duplicate token name '%s' found", tns),
			}
		}
		tks.tokens[tns] = tk

		tk.cred = make(map[string]string)
		for k, v := range cm {
			tk.cred[k] = v.(string)
		}
	}
	return nil
}
