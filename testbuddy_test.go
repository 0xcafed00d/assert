package testbuddy

import (
	"reflect"
	"testing"
)

func test1(a, b, c int) int {
	return a + b + c
}

func test2(a, b, c int) error {
	return nil
}

type dummy struct {
}

var testData = []TestData{
	{test2, Params{1, 2, 3}, Expect{NotNil}},
	{test1, Params{4, "hello", 3}, Expect{9}},
	{test1, Params{1, 2, 3}, Expect{1}},
}

func TestGetFuncName(tst *testing.T) {
	t := (*T)(tst)

	t.Assert(GetShortFuncName(test1)).Equal("test1")
	t.Assert(GetFullFuncName(test1)).Equal("github.com/simulatedsimian/testbuddy.test1")

	t.AssertErr(ConvertTo(5, reflect.TypeOf(test1)))

	err := AutoTest(testData)
	if err != nil {
		tst.Fatal(err)
	}
}
