package retry

import (
	"time"
)

func ExampleCount() {
	done := Count(3 /*times*/, func() bool {
		var shouldGoToNextStep bool
		// :
		if shouldGoToNextStep {
			return true // causes done
		} else {
			// may cause another try
			return false // causes !done after 3-time tries
		}
	})
	if !done {
		// oh no
	}
}

func ExampleWait() {
	done := Wait(3*time.Second, 200*time.Millisecond, func() bool {
		return false
		// impllicit time.Sleep(200*time.Millisecond)
	})
	// about 3 seconds later
	if !done {
	}

	done = Wait(3*time.Second, time.Hour, func() bool {
		return true
	})
	// immediate
	if done {
	}
}

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
