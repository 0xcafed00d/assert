package testbuddy

import (
	"testing"
)

type T struct {
	testing.T
}

type ValSet struct {
	vals []interface{}
	t    *T
}

func (t *T) Verify(vals ...interface{}) *ValSet {
	return &ValSet{vals, t}
}

func (v *ValSet) Expect(expect ...interface{}) *ValSet {

	for i, val := range v.vals {
		_ = i
		_ = val
	}
	return v
}

func (v *ValSet) ExpectError() *ValSet {

	for i, val := range v.vals {
		_ = i
		_ = val
	}
	return v
}

func (v *ValSet) ExpectNoError() *ValSet {

	for i, val := range v.vals {

		_ = i
		_ = val

		if err, isErr := val.(error); isErr {
			_ = err
		}
	}
	return v
}
