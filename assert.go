package assert

import (
	"reflect"
	"testing"
)

type FailFunc func(format string, args ...interface{})

var GetFailFunc func(t *testing.T) FailFunc

func init() {
	GetFailFunc = func(t *testing.T) FailFunc {
		return t.Errorf
	}
}

func True(t *testing.T, val bool) {
	if !val {
		GetFailFunc(t)("Expected: [true] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func False(t *testing.T, val bool) {
	if val {
		GetFailFunc(t)("Expected: [false] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func Equal(t *testing.T, val, expect interface{}) {
	if !reflect.DeepEqual(val, expect) {
		GetFailFunc(t)("Expected: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
	}
}

func NotEqual(t *testing.T, val, expect interface{}) {
	if reflect.DeepEqual(val, expect) {
		GetFailFunc(t)("Not Expecting: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
	}
}

func Nil(t *testing.T, val interface{}) {
	if val != nil && !reflect.ValueOf(val).IsNil() {
		GetFailFunc(t)("Expecting: [nil] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func NotNil(t *testing.T, val interface{}) {
	if val == nil || reflect.ValueOf(val).IsNil() {
		GetFailFunc(t)("Expecting: [not nil] got: [%v]\n%s", val, SourceInfo(2))
	}
}

type TestFunc func(t *testing.T)

func MustPanic(t *testing.T, f TestFunc) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	f(t)
	ci, _ := GetCallerInfo(1)
	GetFailFunc(t)("Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
}
