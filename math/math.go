package math

type Number interface {
	~int | ~uint |
		~int8 | ~uint8 |
		~int16 | ~uint16 |
		~int32 | ~uint32 |
		~int64 | ~uint64 |
		~float32 | ~float64
}

func Abs[T Number](value T) T {
	if value >= 0 {
		return value
	}

	return -value
}
