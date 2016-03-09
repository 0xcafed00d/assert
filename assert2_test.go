package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func boolFunc(a, b bool) (bool, bool) {
	return a, b
}

type TestStruct struct {
	int
}

func TestIsNillable(t *testing.T) {
	var err error

	nillableValues := []interface{}{
		boolFunc,
		&TestStruct{0},
		make(chan int),
		make([]int, 5),
		make(map[string]int),
		err,
		fmt.Errorf("xxx"),
	}

	for _, v := range nillableValues {
		if !isNillable(v) {
			t.Errorf("Type %v is reported as not nillable", reflect.TypeOf(v))
		}
	}

	notNillableValues := []interface{}{
		"string",
		5,
		5.5,
		TestStruct{},
	}

	for _, v := range notNillableValues {
		if isNillable(v) {
			t.Errorf("Type %v is reported as nillable", reflect.TypeOf(v))
		}
	}

}

func TestAssert2(t *testing.T) {
	failCount := 0

	test := Make(t, func(format string, args ...interface{}) {
		failCount++
	})

	failCount = 0
	test(noErrorFunc()).Equal(1, 2)
	if failCount == 1 {
		t.Errorf("Equal(1, 2) failed")
	}

	failCount = 0
	test(noErrorFunc()).Equal(2, 3)
	if failCount == 0 {
		t.Errorf("Equal(2, 3) did not fail")
	}

	failCount = 0
	test(noErrorFunc()).NotEqual(1, 2)
	if failCount == 0 {
		t.Errorf("NotEqual(1, 2) did not fail")
	}

	failCount = 0
	test(noErrorFunc()).NotEqual(2, 3)
	if failCount == 1 {
		t.Errorf("NotEqual(2, 3) failed")
	}

}
