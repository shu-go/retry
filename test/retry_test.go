package test

import (
	"testing"
	"time"

	"bitbucket.org/shu/retry"
)

func TestCount(t *testing.T) {
	a := 0
	done := retry.Count(3, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 3 {
		t.Errorf("have %d want %d", a, 3)
	}

	a = 0
	done = retry.Count(3, func() bool {
		a++
		return true
	})
	if !done {
		t.Errorf("not done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = retry.Count(0, func() bool {
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
	done := retry.Time(8*time.Millisecond, func() bool {
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
	done = retry.Time(8*time.Millisecond, func() bool {
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
	done = retry.Time(3*time.Millisecond, func() bool {
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
	done = retry.Time(3*time.Millisecond, func() bool {
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
	done := retry.Wait(8*time.Millisecond, 3*time.Millisecond, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 3 {
		t.Errorf("have %d want %d", a, 3)
	}

	a = 0
	done = retry.Wait(3*time.Millisecond, 3*time.Millisecond, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}

	a = 0
	done = retry.Wait(3*time.Millisecond, 3*time.Hour, func() bool {
		a++
		return false
	})
	if done {
		t.Errorf("done!?")
	} else if a != 1 {
		t.Errorf("have %d want %d", a, 1)
	}
}
