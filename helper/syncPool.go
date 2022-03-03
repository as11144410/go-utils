package helper

import (
	"fmt"
	"sync"
)

var strPool = sync.Pool{
	New: func() interface{} {
		return "test str"
	},
}

// SyncPool 复用已经使用过的对象,当对象特别大并且使用非常频繁的时候可以大大的减少对象的创建和回收的时间
func SyncPool() {
	str := strPool.Get()
	fmt.Println(str)
	strPool.Put(str)
}
