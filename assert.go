package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, val, expect interface{}) {
	if !reflect.DeepEqual(val, expect) {
		ci, _ := GetCallerInfo(1)
		s := fmt.Sprintf("Expected: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		t.Log(s)
		t.FailNow()
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
