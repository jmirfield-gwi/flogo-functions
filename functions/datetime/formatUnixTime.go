package datetime

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"strconv"
)

type FormatUnixTime struct {

}

func init() {
	function.Register(&FormatUnixTime{})
}

func (s *FormatUnixTime) Name() string {
	return "formatUnixTime"
}

func (s *FormatUnixTime) GetCategory() string {
	return "datetime"
}

func (s *FormatUnixTime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeDateTime}, false
}

func (s *FormatUnixTime) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToDateTime(params[0])
	if err != nil {
		return nil, err
	}
	return strconv.FormatInt(date.Unix(), 10), nil
}