package testbuddy

import (
	"reflect"
	"testing"
)

type T testing.T

type ValTest struct {
	val interface{}
	t   *T
}

func (vt ValTest) Equal(expect interface{}) {
	if !reflect.DeepEqual(vt.val, expect) {
		vt.t.Errorf("Expected: [%v] got: [%v]\n%s", expect, vt.val, SourceInfo(2))
	}
}

func (vt ValTest) NotEqual(expect interface{}) {
	if reflect.DeepEqual(vt.val, expect) {
		ci, _ := GetCallerInfo(1)
		vt.t.Errorf("Not Expecting: [%v] got: [%v]\n%s:[%d]\n%s", expect, vt.val, ci.filename, ci.lineNum, ci.lineSrc)
	}
}

func (t *T) Assert(val interface{}) ValTest {
	return ValTest{val, t}
}

func (t *T) AssertNoErr(val interface{}, err error) ValTest {
	if err != nil {
		ci, _ := GetCallerInfo(1)
		t.Errorf("Expected No Error: got Error: [%v]\n%s:[%d]\n%s", err, ci.filename, ci.lineNum, ci.lineSrc)
	}

	return ValTest{val, t}
}

func (t *T) AssertErr(val interface{}, err error) ValTest {
	if err == nil {
		ci, _ := GetCallerInfo(1)
		t.Errorf("Expected Error: got No Error: [%v]\n%s:[%d]\n%s", err, ci.filename, ci.lineNum, ci.lineSrc)
	}

	return ValTest{val, t}
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
