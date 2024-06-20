package comm_utils

import (
	"path/filepath"
	"runtime"
)

func GetGoBaseDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../../../..")
}
