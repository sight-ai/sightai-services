package comm_utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Int64ArrayToUintArray(in []int64) []uint {
	res := make([]uint, len(in))
	for i, num := range in {
		res[i] = uint(num)
	}
	return res
}

func IntArrayToUintArray(in []int) []uint {
	res := make([]uint, len(in))
	for i, num := range in {
		res[i] = uint(num)
	}
	return res
}

func UintArrayToInt64Array(in []uint) []int64 {
	res := make([]int64, len(in))
	for i, num := range in {
		res[i] = int64(num)
	}
	return res
}

func IntArrayToInt64Array(in []int) []int64 {
	res := make([]int64, len(in))
	for i, num := range in {
		res[i] = int64(num)
	}
	return res
}

func ShuffleUintArray(in []uint) []uint {
	rand.Shuffle(len(in), func(i, j int) { in[i], in[j] = in[j], in[i] })
	return in
}

func GetUintSublistSafe(l []uint, offset, limit uint) []uint {
	if len(l) == 0 {
		return []uint{}
	}

	if limit == 0 {
		return []uint{}
	}

	if int(offset) >= len(l) {
		return []uint{}
	}

	if int(offset+limit) >= len(l) {
		return l[offset:]
	}

	return l[offset : offset+limit]
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RemoveEmptyStringsFromArray(l []string) []string {
	res := []string{}
	for _, s := range l {
		res = append(res, s)
	}
	return res
}
