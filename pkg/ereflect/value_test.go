package ereflect

import (
	"errors"
	"testing"
)

func TestIsNil(t *testing.T) {
	sure(t, IsNil(nil))
	sure(t, !IsNil(1))
}

func TestEqual(t *testing.T) {
	sure(t, Equal(nil, nil))
	sure(t, Equal(1, 1))
	sure(t, Equal("aaa", "aaa"))

	var ch chan struct{}
	sure(t, Equal(nil, ch))
	var m map[string]string
	sure(t, Equal(nil, m))
	var p *int
	sure(t, Equal(nil, p))
	var i interface{}
	sure(t, Equal(nil, i))
	var b []byte
	sure(t, Equal(nil, b))

	sure(t, Equal([]byte{}, []byte{}))
	sure(t, Equal([]byte{0, 1, 2}, []byte{0, 1, 2}))

	sure(t, !Equal(nil, 1))
	sure(t, !Equal([]byte{}, "aaa"))
	sure(t, !Equal(nil, errors.New("mock error")))
}

func TestEqualInteger(t *testing.T) {
	sure(t, EqualInteger(0, 0))
	sure(t, EqualInteger(1, uint(1)))
	sure(t, EqualInteger(uint32(1), int16(1)))
	sure(t, EqualInteger(uint(1), uint8(1)))

	sure(t, !EqualInteger(1, 0))
	sure(t, !EqualInteger(0, "aaa"))
	sure(t, !EqualInteger(-1, uint(0)))
	sure(t, !EqualInteger(uint16(0), int32(-1)))
}

// 为避免package循环引用，写个简单的测试函数
func sure(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Error()
	}
}
