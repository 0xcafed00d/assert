package testbuddy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type CallerInfo struct {
	pc       uintptr
	filename string
	lineNum  int
	lineSrc  string
}

func GetCallerInfo(skip int) (CallerInfo, bool) {
	pc, file, line, ok := runtime.Caller(skip + 1)

	ci := CallerInfo{pc, file, line, "[no source]"}
	if ok {
		data, err := ioutil.ReadFile(file)
		if err == nil {
			lines := bytes.Split(data, []byte{'\n'})
			ci.lineSrc = string(lines[ci.lineNum-1])
		}
	}
	return ci, ok
}

func SourceInfo(skip ...int) string {
	sk := 1
	if len(skip) > 0 {
		sk = skip[0]
	}
	ci, _ := GetCallerInfo(sk)
	return fmt.Sprintf("%s : [%d] : %s", ci.filename, ci.lineNum, ci.lineSrc)
}

func AssertEqual(t *testing.T, val, expect interface{}) {
	if !reflect.DeepEqual(val, expect) {
		ci, _ := GetCallerInfo(1)
		s := fmt.Sprintf("Expected: [%v] got: [%v]\n%s:[%d]\n%s", expect, val, ci.filename, ci.lineNum, ci.lineSrc)
		t.Log(s)
		t.FailNow()
	}
}

type TestingFunc func(t *testing.T)

func MustPanic(t *testing.T, f TestingFunc) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	f(t)

	ci, _ := GetCallerInfo(1)
	s := fmt.Sprintf("Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
	t.Log(s)
	t.FailNow()
}

func MustNotPanic(t *testing.T, f TestingFunc) {
	defer func() {
		if r := recover(); r != nil {
			ci, _ := GetCallerInfo(1)
			s := fmt.Sprintf("Not Expecting Panic:\n%s:[%d]\n%s", ci.filename, ci.lineNum, ci.lineSrc)
			t.Log(s)
			t.FailNow()
		}
	}()
	f(t)
}
