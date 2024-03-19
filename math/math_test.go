package math_test

import (
	"testing"

	"github.com/CP-Algorithm/CpLib/math"
)

type absCase[T math.Number] *struct {
	input    T
	expected T
}

func testAbs[T math.Number](cs []absCase[T]) func(t *testing.T) {
	return func(t *testing.T) {
		for _, c := range cs {
			res := math.Abs(c.input)
			if res != c.expected {
				t.Errorf("testAbs: expected %v, got %v", c.expected, res)
			}
		}
	}
}

func TestAbsFunction(t *testing.T) {
	t.Parallel()

	t.Run("int", testAbs[int]([]absCase[int]{
		{input: 0, expected: 0},
		{input: -12, expected: 12},
		{input: 90, expected: 90},
	}))

	t.Run("uint", testAbs[uint]([]absCase[uint]{
		{input: 679, expected: 679},
		{input: 0, expected: 0},
	}))

	t.Run("int8", testAbs[int8]([]absCase[int8]{
		{input: int8(100), expected: int8(100)},
		{input: int8(0), expected: int8(0)},
		{input: int8(-45), expected: int8(45)},
	}))

	t.Run("uint8", testAbs[uint8]([]absCase[uint8]{
		{input: uint8(0), expected: uint8(0)},
		{input: uint8(40), expected: uint8(40)},
	}))

	t.Run("int16", testAbs[int16]([]absCase[int16]{
		{input: int16(8900), expected: int16(8900)},
		{input: int16(0), expected: int16(0)},
		{input: int16(-7600), expected: int16(7600)},
	}))

	t.Run("uint16", testAbs[uint16]([]absCase[uint16]{
		{input: uint16(0), expected: uint16(0)},
		{input: uint16(10000), expected: uint16(10000)},
	}))

	t.Run("int32", testAbs[int32]([]absCase[int32]{
		{input: int32(8766), expected: int32(8766)},
		{input: int32(0), expected: int32(0)},
		{input: int32(-0x123), expected: int32(0x123)},
	}))

	t.Run("uint32", testAbs[uint32]([]absCase[uint32]{
		{input: uint32(0o0), expected: uint32(0o0)},
		{input: uint32(0b10101), expected: uint32(0b10101)},
	}))

	t.Run("int64", testAbs[int64]([]absCase[int64]{
		{input: int64(500000000000), expected: int64(500000000000)},
		{input: int64(0), expected: int64(0)},
		{input: int64(-0), expected: int64(0)},
		{input: int64(-0xABC), expected: int64(0xABC)},
	}))

	t.Run("uint64", testAbs[uint64]([]absCase[uint64]{
		{input: uint64(0), expected: uint64(0)},
		{input: uint64(89000037633), expected: uint64(89000037633)},
	}))

	t.Run("float32", testAbs[float32]([]absCase[float32]{
		{input: float32(0.1), expected: float32(0.1)},
		{input: float32(-0.0), expected: float32(0.0)},
		{input: float32(0), expected: float32(0)},
		{input: float32(-123.5), expected: float32(123.5)},
	}))

	t.Run("float64", testAbs[float64]([]absCase[float64]{
		{input: float64(0.19238), expected: float64(0.19238)},
		{input: float64(-0.0), expected: float64(0.0)},
		{input: float64(0), expected: float64(0)},
		{input: float64(-100000.67), expected: float64(100000.67)},
	}))
}
