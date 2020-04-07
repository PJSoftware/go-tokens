package gotokens

import "fmt"

// Token represents an individual token set extracted from Tokens
type Token struct {
	cred map[string]string
}

// Credential returns the value of the named credential from the token
func (tk *Token) Credential(cname string) (string, error) {
	if cv, ok := tk.cred[cname]; ok {
		return cv, nil
	}

	return "", &Error{
		Code:    EBADCREDENTIAL,
		Message: fmt.Sprintf("Unknown credential '%s'", cname),
	}
}
