package assert

import "testing"

func Test1(t *testing.T) {
	test := Make(t)

	test(true).
		Equal(true, false).
		NotEqual(true).
		HasError().
		NoError().IsNil()
}
