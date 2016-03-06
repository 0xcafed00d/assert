package assert

import (
	"reflect"
	"testing"
)

type failFunc func(format string, args ...interface{})

type Results struct {
	results []interface{}
	onFail  failFunc
}

type DoTestFunc func(args ...interface{}) *Results

func Make(t *testing.T, f ...failFunc) DoTestFunc {
	onFail := t.Errorf

	if len(f) > 0 {
		onFail = f[0]
	}

	return func(args ...interface{}) *Results {
		return &Results{args, onFail}
	}
}

func (r *Results) Equal(expect ...interface{}) *Results {
	if !reflect.DeepEqual(r.results, expect) {
		r.onFail("Equal Expected: [%v] got: [%v]\n%s", expect, r.results, SourceInfo(2))
	}
	return r
}

func (r *Results) NotEqual(expect ...interface{}) *Results {
	if reflect.DeepEqual(r.results, expect) {
		r.onFail("NotEqual Not Expecting: [%v] got: [%v]\n%s", expect, r.results, SourceInfo(2))
	}
	return r
}
