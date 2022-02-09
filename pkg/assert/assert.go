package assert

import "github.com/Etuloser/ego/pkg/ereflect"

// TestingT 单元测试中的 *testing.T 和 *testing.B 都满足该接口
type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper interface {
	Helper()
}

func Equal(t TestingT, expected interface{}, actual interface{}, msg ...string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if !ereflect.Equal(expected, actual) {
		t.Errorf("%s expected=%+v, actual=%+v", msg, expected, actual)
	}
}

// IsNotNil 比如有时我们需要对 error 类型不等于 nil 做断言，但是我们并不关心 error 的具体值是什么
func IsNotNil(t TestingT, actual interface{}, msg ...string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if ereflect.IsNil(actual) {
		t.Errorf("%s expected not nil, but actual=%+v", msg, actual)
	}
}
