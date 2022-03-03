package helper

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//attempts：最多重试次数；
//sleep：调用失败后的等待时间；
//fn：重试的函数。函数的类型是func() error，如果你的重试函数定义并不是这样，可以通过闭包包一下。
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}
		if attempts--; attempts > 0 {
			fmt.Errorf("retry func error: %s. attemps #%d after %s.", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func NoRetryError(err error) stop {
	return stop{err}
}

// 实现2
type RetryInterface interface {
	Run() error       // 实际运行的方法
	IsExpected() bool // 确认方法返回结果是否符合预期
}

func Retry2(f RetryInterface, maxRetryTimes int) error {
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxRetryTimes; i++ {
		err := f.Run()
		if err != nil {
			fmt.Printf("%s", err.Error())
			switch err.Error() {
			case "time out":
				time.Sleep(time.Duration(rand.Int31n(100)))
			case "params error":
				return errors.New("params error")
			}
			continue
		}
		if err := recover(); err != nil {
			fmt.Printf("%s", err)
			time.Sleep(time.Duration(rand.Int31n(100)))
			continue
		}
		if f.IsExpected() {
			return nil
		}
	}
	return errors.New("Maximum retry attempts was exceeded.")
}
