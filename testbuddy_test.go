package testbuddy

import (
	"reflect"
	"testing"
)

func test1() {

}

func TestGetFuncName(tst *testing.T) {
	t := (*T)(tst)

	t.Assert(GetShortFuncName(test1)).Equal("test1")
	t.Assert(GetFullFuncName(test1)).Equal("github.com/simulatedsimian/testbuddy.test1")

	t.AssertNoErr(ConvertTo(5, reflect.TypeOf(test1)))

}
