package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	function.ResolveAliases()
}

func TestSuccessfulFormatUnixTime_Eval(t *testing.T) {
	var in = &FormatUnixTime{}
	final, err:= in.Eval("2022-06-19T19:48:37.073976111Z")
	assert.Nil(t, err)
	assert.Equal(t, "1655668117", final)
}

func TestFailingFormatTooShort_Eval(t *testing.T) {
	var in = &FormatUnixTime{}
	final, err:= in.Eval("123")
	assert.Nil(t, final)
	assert.Equal(t, err.Error(), "parse [123] to time error: unrecognized format, too short 123")
}

func TestFailingInvalidCharacters_Eval(t *testing.T) {
	var in = &FormatUnixTime{}
	final, err:= in.Eval("abc")
	assert.Nil(t, final)
	assert.Equal(t, err.Error(), "parse [abc] to time error: Could not find format for \"abc\"")
}