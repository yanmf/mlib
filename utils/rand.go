package utils

import (
	//"errors"
	//"time"

	"math/rand"
	"time"
	//"util/log"
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Rand(args ...int32) int32 {
	return RandSide(globalRand, args...)
}

func Rand1w() int32 {
	return Rand(10000)
}

func Rand1wOK(val int32) bool {
	return val > Rand1w()
}

func Randf32() float32 {
	return rand.Float32()
}

func Randf64() float64 {
	return rand.Float64()
}

func RandSide(randSide *rand.Rand, args ...int32) int32 {
	var l = len(args)
	switch l {
	case 0:
		return randSide.Int31()
	case 1:
		max := args[0]
		if max <= 0 {
			return 0
		}
		return randSide.Int31n(max)
	default:
		min, max := args[0], args[1]
		if min == max {
			return min
		}
		if max < min {
			min, max = max, min
		}
		return min + randSide.Int31n(max-min)
	}
}
