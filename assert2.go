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

func isNillable(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr
}

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

func (r *Results) NoError() *Results {
	for _, v := range r.results {
		if err, ok := v.(error); ok {
			r.onFail("NoError Expecting: [no error] got error: [%v]\n%s", err, SourceInfo(2))
		}
	}
	return r
}

func (r *Results) HasError() *Results {
	for _, v := range r.results {
		if _, ok := v.(error); ok {
			return r
		}
	}
	r.onFail("NoError Expecting: [error] got no error:\n%s", SourceInfo(2))
	return r
}

func (r *Results) IsNil() *Results {
	for _, val := range r.results {
		if isNillable(reflect.TypeOf(val)) && !reflect.ValueOf(val).IsNil() {
			r.onFail("IsNil Expecting: [nil] got: [%v]\n%s", val, SourceInfo(2))
		}
	}
	return r
}

func (r *Results) NotNil() *Results {
	for _, val := range r.results {
		if val == nil || reflect.ValueOf(val).IsNil() {
			r.onFail("NotNil Expecting: [not nil] got: [%v]\n%s", val, SourceInfo(2))
		}
	}
	return r
}
