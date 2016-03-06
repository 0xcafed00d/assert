package assert

import "testing"

func Test1(t *testing.T) {
	test := Make(t)

	test(true).
		Equal(true).
		NotEqual(true)
}
