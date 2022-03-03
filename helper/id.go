package helper

import (
	idWorker "github.com/gitstliu/go-id-worker"
	"strconv"
)

// MakeUniqueSn 雪花生成唯一
func MakeUniqueSn(prefix string) string {
	worker := &idWorker.IdWorker{}
	err := worker.InitIdWorker(1000000000000000, 1)
	if err != nil {
		return ""
	}
	newId, newIdErr := worker.NextId()
	if newIdErr != nil {
		return ""
	}
	return prefix + strconv.FormatInt(newId, 10)
}
