package testbuddy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
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

func CallFunction(f interface{}, args []interface{}) ([]interface{}, error) {
	fval := reflect.ValueOf(f)
	argVals := []reflect.Value{}

	numIn := fval.Type().NumIn()

	if len(args) != fval.Type().NumIn() {
		return nil, fmt.Errorf("Incorrect Number of Args. Expected: %d Got: %d", numIn, len(args))
	}

	for i := 0; i < numIn; i++ {
		arg, err := ConvertTo(args[i], fval.Type().In(i))
		if err != nil {
			return nil, fmt.Errorf("Param #%d Error: [%v]", i, err)
		}
		argVals = append(argVals, reflect.ValueOf(arg))
	}

	retVals := fval.Call(argVals)
	var ret []interface{}

	for i := 0; i < len(retVals); i++ {
		ret = append(ret, retVals[i].Interface())
	}

	return ret, nil
}

func ConvertTo(i interface{}, to reflect.Type) (interface{}, error) {
	from := reflect.TypeOf(i)
	if from.ConvertibleTo(to) {
		return reflect.ValueOf(i).Convert(to).Interface(), nil
	}
	return nil, fmt.Errorf("Cannot Convert From %v to %v", from, to)
}

func GetFullFuncName(i interface{}) string {
	p := reflect.ValueOf(i).Pointer()
	if p == 0 {
		return "nil"
	}
	return runtime.FuncForPC(p).Name()
}

func GetShortFuncName(i interface{}) string {
	name := GetFullFuncName(i)
	return name[strings.LastIndex(name, ".")+1:]
}

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
	return fmt.Sprintf("%s : [%d]\n%s", ci.filename, ci.lineNum, strings.Trim(ci.lineSrc, " \t"))
}
