package gtc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A shorthand where, if text == "", we assert that there was
// no error... and otherwise, we assert that an error occurred,
// and verify that err.Error() == text.
//
// The `fail` bool determines whether an unexpected result should
// call t.FailNow(), rather than just t.Error().
func AssertError(t *testing.T, err error, text string, fail bool) {
	if !assertError(t, err, text) && fail {
		t.FailNow()
	}
}

// Underlying function for AssertError - does assertions and
// returns success.
func assertError(t *testing.T, err error, text string) bool {
	if text == "" {
		return assert.NoError(t, err)
	} else {
		if assert.Error(t, err) {
			return assert.Equal(t, text, err.Error())
		} else {
			return false
		}
	}
}
