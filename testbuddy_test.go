package testbuddy

import (
	"reflect"
	"testing"
)

func test1(a, b, c int) int {
	return a + b + c
}

var testData = []TestData{
	{test1, Params{1, 2, 3}, Expect{6}},
	{test1, Params{4, 2, 3}, Expect{9}},
	{test1, Params{1, 2, 3}, Expect{1}},
}

func TestGetFuncName(tst *testing.T) {
	t := (*T)(tst)

	t.Assert(GetShortFuncName(test1)).Equal("test1")
	t.Assert(GetFullFuncName(test1)).Equal("github.com/simulatedsimian/testbuddy.test1")

	t.AssertErr(ConvertTo(5, reflect.TypeOf(test1)))

	params := []interface{}{1, 2, 3}

	ret, err := CallFunction(test1, params)

	if err != nil {
		tst.Fatalf("%s", err)
	}

	t.Assert(ret[0]).Equal(6)

	AutoTest(tst, testData)
}
