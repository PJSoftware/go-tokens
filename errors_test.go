package gotokens_test

import (
	"fmt"
	"testing"

	"github.com/pjsoftware/gotokens"
)

func TestErrorTypes(t *testing.T) {
	testErrorOutput(t, nil, gotokens.ENOERROR,
		"", "", "")

	err1 := fmt.Errorf("Simple error")
	testErrorOutput(t, err1, gotokens.EINTERNAL,
		"An internal error has occurred: Simple error", "NoContext", "Simple error")

	err2 := fmt.Errorf("Extended error: %v", err1)
	testErrorOutput(t, err2, gotokens.EINTERNAL,
		"An internal error has occurred: Extended error: Simple error", "NoContext", "Extended error: Simple error")

	gte1 := &gotokens.Error{
		Code:    gotokens.ENOSEARCHPATH,
		Message: "Path not found",
	}
	testErrorOutput(t, gte1, gotokens.ENOSEARCHPATH,
		"Path not found", "NoContext", "<E_NO_SEARCH_PATH> Path not found")

	gte2 := &gotokens.Error{Op: "Testing", Err: gte1}
	testErrorOutput(t, gte2, gotokens.ENOSEARCHPATH,
		"Path not found", "NoContext", "Testing: <E_NO_SEARCH_PATH> Path not found")

	gte3 := &gotokens.Error{Op: "StillTesting", Context: "SomeContext", Err: gte2}
	testErrorOutput(t, gte3, gotokens.ENOSEARCHPATH,
		"Path not found", "SomeContext", "StillTesting/SomeContext: Testing: <E_NO_SEARCH_PATH> Path not found")
}

func testErrorOutput(t *testing.T, err error, expc, expm, expx, exps string) {
	if got := gotokens.ErrorCode(err); got != expc {
		t.Errorf("Expected Code '%s' but got '%s'", expc, got)
	}
	if got := gotokens.ErrorMessage(err); got != expm {
		t.Errorf("Expected Message '%s' but got '%s'", expm, got)
	}
	if got := gotokens.ErrorContext(err); got != expx {
		t.Errorf("Expected Context '%s' but got '%s'", expx, got)
	}
	if err != nil {
		if got := err.Error(); got != exps {
			t.Errorf("Expected String '%s' but got '%s'", exps, got)
		}
	}
}

func testErrorCode(t *testing.T, err error, want string) {
	if err == nil {
		t.Errorf("Expected error '%s' but none occurred", want)
	}
	got := gotokens.ErrorCode(err)
	if got != want {
		t.Errorf("Expected error code %s; got %s", want, err)
	}
}

func testErrorContext(t *testing.T, err error, want string) {
	got := gotokens.ErrorContext(err)
	if got != want {
		t.Errorf("Expected error context '%s'; got '%s': %s", want, got, err)
	}
}
