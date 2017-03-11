# retry : retry routines #

[![GoDoc](https://godoc.org/bitbucket.org/shu/retry?status.svg)](https://godoc.org/bitbucket.org/shu/retry)

## Install ##

```
go get bitbucket.org/shu/retry
```

## How to use ##

Count

```
#!go

done := retry.Count(3 /*times*/, func() bool {
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
```

Wait

```
#!go

done := retry.Wait(3*time.Second, 200*time.Millisecond, func() bool {
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
```

## License ##

MIT.
