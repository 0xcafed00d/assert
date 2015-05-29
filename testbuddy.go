package testbuddy

import (
	"reflect"
	"testing"
)

type T testing.T

func (t *T) AssertEqual(val interface{}) func(expect interface{}) {

	return func(expect interface{}) {
		if !reflect.DeepEqual(val, expect) {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Expected: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		}
	}
}

func (t *T) AssertEqualsNoErr(val interface{}, err error) func(expect interface{}) {
	if err != nil {
		ci, _ := GetCallerInfo(1)
		t.Errorf("Expected No Error: got Error: [%v]\n%s:[%d]\n%s", err, ci.filename, ci.lineNum, ci.lineSrc)
	}

	return func(expect interface{}) {
		if !reflect.DeepEqual(val, expect) {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Expected: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		}
	}
}

func (t *T) AssertEqualsErr(val interface{}, err error) func(expect interface{}) {
	if err == nil {
		ci, _ := GetCallerInfo(1)
		t.Errorf("Expected Error: got No Error: [%v]\n%s:[%d]\n%s", err, ci.filename, ci.lineNum, ci.lineSrc)
	}

	return func(expect interface{}) {
		if !reflect.DeepEqual(val, expect) {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Expected: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		}
	}
}

func (t *T) AssertNotEqualsErr(val interface{}, err error) func(expect interface{}) {
	if err != nil {
		ci, _ := GetCallerInfo(1)
		t.Errorf("Expected Error: got: [%v]\n%s:[%d]\n%s", err, ci.filename, ci.lineNum, ci.lineSrc)
	}

	return func(expect interface{}) {
		if reflect.DeepEqual(val, expect) {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Not Expecting: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		}
	}
}

type TestingFunc func(t *T)

func (t *T) MustPanic(f TestingFunc) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	f(t)
	ci, _ := GetCallerInfo(1)
	t.Errorf("Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
}

func (t *T) MustNotPanic(f TestingFunc) {
	defer func() {
		if r := recover(); r != nil {
			ci, _ := GetCallerInfo(1)
			t.Errorf("Not Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
		}
	}()
	f(t)
}
