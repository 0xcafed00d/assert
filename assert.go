package assert

import (
	"reflect"
	"testing"
)

var FailFunc func(format string, args ...interface{})

func getFailFunc(t *testing.T) func(format string, args ...interface{}) {
	if FailFunc == nil {
		return t.Errorf
	}
	return FailFunc
}

func True(t *testing.T, val bool) bool {
	if !val {
		getFailFunc(t)("Expected: [true] got: [%v]\n%s", val, SourceInfo(2))
		return true
	}
	return false
}

func False(t *testing.T, val bool) bool {
	if val {
		getFailFunc(t)("Expected: [false] got: [%v]\n%s", val, SourceInfo(2))
		return true
	}
	return true
}

func Equal(t *testing.T, val, expect interface{}) bool {
	if !reflect.DeepEqual(val, expect) {
		getFailFunc(t)("Expected: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
		return true
	}
	return false
}

func NotEqual(t *testing.T, val, expect interface{}) bool {
	if reflect.DeepEqual(val, expect) {
		getFailFunc(t)("Not Expecting: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
		return true
	}
	return false
}

func Nil(t *testing.T, val interface{}) bool {
	if val != nil && !reflect.ValueOf(val).IsNil() {
		getFailFunc(t)("Expecting: [nil] got: [%v]\n%s", val, SourceInfo(2))
		return true
	}
	return false
}

func NotNil(t *testing.T, val interface{}) bool {
	if val == nil || reflect.ValueOf(val).IsNil() {
		getFailFunc(t)("Expecting: [not nil] got: [%v]\n%s", val, SourceInfo(2))
		return true
	}
	return false
}

type TestFunc func(t *testing.T)

func MustPanic(t *testing.T, f TestFunc) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	f(t)
	ci, _ := GetCallerInfo(1)
	getFailFunc(t)("Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
}

func MustNotPanic(t *testing.T, f TestFunc) {
	defer func() {
		if r := recover(); r != nil {
			ci, _ := GetCallerInfo(1)
			getFailFunc(t)("Not Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
		}
	}()
	f(t)
}
