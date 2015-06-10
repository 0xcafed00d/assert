package testbuddy

import (
	"fmt"
	"reflect"
)

type Params []interface{}
type Expect []interface{}

type StructNotNil struct {
}

var NotNil = StructNotNil{}

type TestData struct {
	F interface{}
	P Params
	E Expect
}

func AutoTest(data []TestData) error {

	makeErr := func(i int, e error, info string) error {
		return fmt.Errorf("Test #%d error: [%v]\n%s", i, e, info)
	}

	for i, tst := range data {
		results, err := CallFunction(tst.F, tst.P)
		if err != nil {
			return makeErr(i, err, SourceInfo(2))
		}

		if len(results) != len(tst.E) {
			err := fmt.Errorf("Incorrect Returned Value Count. Expected: %d Got: %d", len(tst.E), len(results))
			return makeErr(i, err, SourceInfo(2))
		}

		for i, res := range results {

			if _, isNotNil := tst.E[i].(StructNotNil); isNotNil {
				if res == nil {
					err = fmt.Errorf("Returned Value #%d error: Expected <not nil> Got <nil>", i)
					return makeErr(i, err, SourceInfo(2))
				}
			} else {
				r, err := ConvertTo(res, reflect.TypeOf(tst.E[i]))
				if err != nil {
					err = fmt.Errorf("Returned Value #%d error: [%v]", i, err)
					return makeErr(i, err, SourceInfo(2))
				}

				if !reflect.DeepEqual(r, tst.E[i]) {
					err = fmt.Errorf("Returned Value #%d error: Expected %v Got %v", i, tst.E[i], r)
					return makeErr(i, err, SourceInfo(2))
				}
			}

		}
	}

	return nil
}
