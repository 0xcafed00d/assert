package assert

import (
	"errors"
	"testing"
)

var failCalled = false
var testVar int
var testNilPtr *int

func errorFunc() (int, error, int) {
	return 1, errors.New("test error"), 9
}

func nilErrorFunc() (int, error, int) {
	var err error

	return 1, err, 9
}

func noErrorFunc() (int, int) {
	return 1, 2
}

func TestAssert(t *testing.T) {
	GetFailFunc = func(t *testing.T) FailFunc {
		return func(format string, args ...interface{}) {
			failCalled = true
		}
	}

	// -- test True ------------------------------------------------
	failCalled = false
	True(t, false)
	if !failCalled {
		t.Errorf("assert.True(false) did not fail")
	}

	failCalled = false
	True(t, true)
	if failCalled {
		t.Errorf("assert.True(true) incorrectly failed")
	}

	// -- test False ------------------------------------------------
	failCalled = false
	False(t, true)
	if !failCalled {
		t.Errorf("assert.False(true) did not fail")
	}

	failCalled = false
	False(t, false)
	if failCalled {
		t.Errorf("assert.False(false) incorrectly failed")
	}

	// -- test Equal ------------------------------------------------
	failCalled = false
	Equal(t, 1, 2)
	if !failCalled {
		t.Errorf("assert.Equal(1,2) did not fail")
	}

	failCalled = false
	Equal(t, 1, 1)
	if failCalled {
		t.Errorf("assert.Equal(1,1) incorrectly failed")
	}

	// -- test NotEqual ------------------------------------------------
	failCalled = false
	NotEqual(t, 1, 1)
	if !failCalled {
		t.Errorf("assert.Equal(1,1) did not fail")
	}

	failCalled = false
	NotEqual(t, 1, 2)
	if failCalled {
		t.Errorf("assert.Equal(1,2) incorrectly failed")
	}

	// -- test Nil ------------------------------------------------
	failCalled = false
	Nil(t, &testVar)
	if !failCalled {
		t.Errorf("assert.Nil(&testVar) did not fail")
	}

	failCalled = false
	Nil(t, testNilPtr)
	if failCalled {
		t.Errorf("assert.Nil(testNilPtr) incorrectly failed")
	}

	failCalled = false
	Nil(t, nil) // untyped nil
	if failCalled {
		t.Errorf("assert.Nil(nil)  incorrectly failed")
	}

	/* FIXME panics
	failCalled = false
	Nil(t, 1)
	if failCalled {
		t.Errorf("assert.Nil(nil) did not fail")
	}
	*/

	// -- test NotNil ------------------------------------------------
	failCalled = false
	NotNil(t, nil) // untyped nil
	if !failCalled {
		t.Errorf("assert.NotNil(nil) did not fail")
	}

	failCalled = false
	NotNil(t, testNilPtr)
	if !failCalled {
		t.Errorf("assert.NotNil(testNilPtr) did not fail")
	}

	failCalled = false
	NotNil(t, &testVar)
	if failCalled {
		t.Errorf("assert.NotNil(&testVar) incorrectly failed")
	}

	// -- test MustPanic ------------------------------------------------
	failCalled = false
	MustPanic(t, func(t *testing.T) {
		panic(1)
	})
	if failCalled {
		t.Errorf("MustPanic incorreclty failed")
	}

	failCalled = false
	MustPanic(t, func(t *testing.T) {
	})
	if !failCalled {
		t.Errorf("MustPanic did not fail")
	}

	// -- test Error detedction ------------------------------------------------
	failCalled = false
	NoError(t, Pack(errorFunc()))
	if !failCalled {
		t.Errorf("NoError: did not fail on error")
	}

	failCalled = false
	NoError(t, Pack(nilErrorFunc()))
	if failCalled {
		t.Errorf("NoError: failed on nil error")
	}

	failCalled = false
	NoError(t, Pack(noErrorFunc()))
	if failCalled {
		t.Errorf("NoError: failed on no error")
	}

	failCalled = false
	HasError(t, Pack(errorFunc()))
	if failCalled {
		t.Errorf("HasError: failed on error")
	}

	failCalled = false
	HasError(t, Pack(nilErrorFunc()))
	if !failCalled {
		t.Errorf("HasError: failed on nil error")
	}

	failCalled = false
	HasError(t, Pack(noErrorFunc()))
	if !failCalled {
		t.Errorf("HasError: failed on no error")
	}

}
