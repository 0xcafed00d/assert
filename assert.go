package testbuddy

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
