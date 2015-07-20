package assert

import (
	"errors"
	"reflect"
	"testing"
)

func test1(a, b, c int) int {
	return a + b + c
}

func test2(a, b, c int) error {
	return nil
}

func test3(vals ...interface{}) {

}

func test4() (int, string, error) {
	return 1, "1", errors.New("1")
}

type dummy struct {
}

var testData = []TestData{
	{test2, Params{1, 2, 3}, Expect{NotNil}},
	{test1, Params{4, "hello", 3}, Expect{9}},
	{test1, Params{1, 2, 3}, Expect{1}},
}

func TestGetFuncName(t *testing.T) {
	tt := T{t}

	tt.Assert(GetShortFuncName(test1)).Equal("test1")
	tt.Assert(GetFullFuncName(test1)).Equal("github.com/simulatedsimian/testbuddy.test1")

	tt.AssertErr(ConvertTo(5, reflect.TypeOf(test1)))

	err := AutoTest(testData)
	if err != nil {
		tt.Fatal(err)
	}

	test3(test4())
}

func TestVerify(t *testing.T) {

}
