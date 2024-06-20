package mysql

import "strings"

const (
	GormErrorCodeDuplicateEntry = "1062"
)

// example error:
// Error 1062: Duplicate entry '3-1' for key 'device_watch_story.idx_device_id_story_id'
func GetGormErrorCode(err error) string {
	ss := strings.Split(err.Error(), " ")
	if len(ss) < 2 {
		return ""
	}
	code := ss[1]
	return code[:len(code)-1]
}

type IDStruct struct {
	ID uint
}

func IDStructArrayToUintArray(ss []*IDStruct) []uint {
	res := make([]uint, len(ss))
	for i, s := range ss {
		res[i] = s.ID
	}
	return res
}

func IDStructArrayToInt64Array(ss []*IDStruct) []int64 {
	res := make([]int64, len(ss))
	for i, s := range ss {
		res[i] = int64(s.ID)
	}
	return res
}
