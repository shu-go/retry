package retry

import (
	"testing"
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

func TestCount(t *testing.T) {
	a := 0
	done := Count(3, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 3 {
		t.Errorf("have %d want %d", a, 3)
	}

	a = 0
	done = Count(3, func() bool {
		a++
		return true
	})
	if !done {
		t.Errorf("not done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = Count(0, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 0 {
		t.Errorf("have %d want %d", a, 0)
	}
}

func TestTime(t *testing.T) {
	a := 0
	done := Time(8*time.Millisecond, func() bool {
		a++
		time.Sleep(3 * time.Millisecond)
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 3 {
		t.Errorf("have %d want %d", a, 3)
	}

	a = 0
	done = Time(8*time.Millisecond, func() bool {
		a++
		time.Sleep(3 * time.Millisecond)
		return true
	})
	if !done {
		t.Errorf("not done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = Time(3*time.Millisecond, func() bool {
		a++
		time.Sleep(3 * time.Millisecond)
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = Time(3*time.Millisecond, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a == 0 {
		t.Errorf("have %s want %d", "so many", a)
	}
}

func TestWait(t *testing.T) {
	a := 0
	done := Wait(8*time.Millisecond, 3*time.Millisecond, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 3 {
		t.Errorf("have %d want %d", a, 3)
	}

	a = 0
	done = Wait(3*time.Millisecond, 3*time.Millisecond, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = Wait(3*time.Millisecond, 3*time.Hour, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}
}
