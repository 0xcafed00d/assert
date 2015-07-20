package assert

import (
	"reflect"
	"testing"
)

func True(t *testing.T, val bool) {
	if !val {
		t.Errorf("Expected: [true] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func False(t *testing.T, val bool) {
	if val {
		t.Errorf("Expected: [false] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func Equal(t *testing.T, val, expect interface{}) {
	if !reflect.DeepEqual(val, expect) {
		t.Errorf("Expected: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
	}
}

func NotEqual(t *testing.T, val, expect interface{}) {
	if reflect.DeepEqual(val, expect) {
		t.Errorf("Not Expecting: [%v] got: [%v]\n%s", expect, val, SourceInfo(2))
	}
}

func Nil(t *testing.T, val interface{}) {
	if val != nil && !reflect.ValueOf(val).IsNil() {
		t.Errorf("Expecting: [nil] got: [%v]\n%s", val, SourceInfo(2))
	}
}

func NotNil(t *testing.T, val interface{}) {
	if val == nil || reflect.ValueOf(val).IsNil() {
		t.Errorf("Expecting: [not nil] got: [%v]\n%s", val, SourceInfo(2))
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
	t.Errorf("Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
}

func MustNotPanic(t *testing.T, f TestFunc) {
	defer func() {
		if r := recover(); r != nil {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Not Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
		}
	}()
	f(t)
}
