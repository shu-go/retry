package retry

import (
	"time"
)

type Func func() (done bool)

func Count(count int, f Func) bool {
	for i := 1; i <= count; i++ {
		done := f()
		if done {
			return true
		}
	}
	return false
}

func Time(timeout time.Duration, f Func) bool {
	return Wait(timeout, 0, f)
}

func Wait(timeout, wait time.Duration, f Func) bool {
	start := time.Now()
	for {
		done := f()
		if done {
			return true
		}
		if timeout <= time.Since(start)+wait {
			break
		}
		time.Sleep(wait)
	}
	return false
}
