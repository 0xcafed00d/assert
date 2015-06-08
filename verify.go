package testbuddy

type ValSet struct {
	val interface{}
	t   *T
}

func (t *T) Verify(vals ...interface{}) *ValSet {
	return &ValTest{val, t}
}
