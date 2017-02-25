// retry routines
package retry

import (
	"time"
)

// Func is a function type that is passed to retry routines.
// If it returns true, retry routines end immediately.
type Func func() (done bool)

// Count retries the function f at most count times.
func Count(count int, f Func) bool {
	for i := 1; i <= count; i++ {
		done := f()
		if done {
			return true
		}
	}
	return false
}

// Time retries the function f until the duration timeout is passed.
func Time(timeout time.Duration, f Func) bool {
	return Wait(timeout, 0, f)
}

// Wait retries the function f until the duration timeout is passed.
// It also sleeps for the duration wait between each try.
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
