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

type Ignore struct {
}

type DoTestFunc func(args ...interface{}) *Results

func isNillable(v interface{}) bool {
	if v == nil {
		return true
	}

	k := reflect.TypeOf(v).Kind()
	return k == reflect.Ptr || k == reflect.Chan || k == reflect.Func ||
		k == reflect.Interface || k == reflect.Map || k == reflect.Slice
}

func isNil(val interface{}) bool {
	return isNillable(val) && (val == nil || reflect.ValueOf(val).IsNil())
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

	if len(r.results) != len(expect) {
		r.onFail("Equal Failed with Parameter count mismatch expected: [%v] got: [%v]\n%s", expect, r.results, SourceInfo(2))
	}

	for i := range r.results {
		if reflect.TypeOf(expect[i]) == reflect.TypeOf(Ignore{}) {
			break
		}

		// if return value is a pointer then derefernce it to the value before comparison
		if !isNil(r.results[i]) && reflect.TypeOf(r.results[i]).Kind() == reflect.Ptr {
			r.results[i] = reflect.ValueOf(r.results[i]).Elem().Interface()
		}

		// if expect value is a pointer then derefernce it to the value before comparison
		if !isNil(expect[i]) && reflect.TypeOf(expect[i]).Kind() == reflect.Ptr {
			expect[i] = reflect.ValueOf(expect[i]).Elem().Interface()
		}

		if !reflect.DeepEqual(r.results[i], expect[i]) {
			r.onFail("Equal Expected: [%v] got: [%v]\n%s", expect, r.results, SourceInfo(2))
		}
	}
	return r
}

func (r *Results) NotEqual(expect ...interface{}) *Results {

	if len(r.results) != len(expect) {
		r.onFail("Not Equal Failed with Parameter count mismatch expected: [%v] got: [%v]\n%s", expect, r.results, SourceInfo(2))
	}

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
		if val != nil && isNillable(val) && !reflect.ValueOf(val).IsNil() {
			r.onFail("IsNil Expecting: [nil] got: [%v]\n%s", val, SourceInfo(2))
		}
	}
	return r
}

func (r *Results) NotNil() *Results {
	for _, val := range r.results {
		if val == nil || (isNillable(val) && reflect.ValueOf(val).IsNil()) {
			r.onFail("NotNil Expecting: [not nil] got: [%v]\n%s", val, SourceInfo(2))
		}
	}
	return r
}
