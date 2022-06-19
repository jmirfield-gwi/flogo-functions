package datetime

import (
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	function.ResolveAliases()
}

func TestFnFormatUnixTime_Eval(t *testing.T) {
	var in = &FormatUnixTime{}
	final, err:= in.Eval("2022-06-19T19:48:37.073976111Z")
	assert.Nil(t, err)
	assert.Equal(t, "1655668117", final)
}