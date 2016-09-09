package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v "./validator"
)

type toErrOrNotToErr func(assert.TestingT, error, ...interface{}) bool

var Error = assert.Error
var NoError = assert.NoError

func TestValidator(t *testing.T) {
	type testdata struct {
		value   string
		c       bool
		errFunc toErrOrNotToErr
	}

	tests := []testdata{
		{"", false, Error},
		{"     ", true, Error},
		{"ab", false, Error},
		{"    ab", true, Error},
		{"ab    ", true, Error},
		{"   ab ", true, Error},
		{"abc", false, NoError},
		{"abcd", false, NoError},
		{" abc", true, NoError},
		{"abc ", true, NoError},
		{"a b c", false, Error},
		{"abcdefg", false, NoError},
	}

	for _, test := range tests {
		n := &v.Name{}
		n.Name = test.value
		c, err := n.Validate()
		assert.Equal(t, c, test.c)
		test.errFunc(t, err)
	}
}

func TestValidatorMaps(t *testing.T) {
	type testdata struct {
		key      string
		value    string
		c        bool
		errFunc  toErrOrNotToErr
		length   int
		shouldBe map[string]string
	}

	tests := []testdata{
		{"", "", false, NoError, 0, nil},
		{"hello", "", false, NoError, 0, nil},
		{"hello", "bye", false, NoError, 1, map[string]string{"hello": "bye"}},
		{"    hell ", "   bye   ", false, NoError, 1, map[string]string{"hell": "bye"}},
	}

	for _, test := range tests {
		n := &v.Tags{}
		n.Tags = map[string]string{test.key: test.value}
		c, err := n.Validate()
		assert.Equal(t, c, test.c)
		test.errFunc(t, err)
		assert.Len(t, n.Tags, test.length)

		for k, v := range test.shouldBe {
			assert.Equal(t, n.Tags[k], v)
		}
	}

}
