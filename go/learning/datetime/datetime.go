package main

import (
	"fmt"
	"time"
)

func main() {
	curTime := time.Now()
	fmt.Printf("curTime=%v, curTime=%T\n", curTime, curTime)

	// The format strings time.ANSIC, time.Kitchen, time.Layout, time.Kitchen
	// are untyped constants defined in time package
	fmt.Printf("curTime in Kitchen format=%v\n", curTime.Format(time.Kitchen))
	fmt.Printf("curTime in ANSIC format=%v\n", curTime.Format(time.ANSIC))
	fmt.Printf("curTime in Layout format=%v\n", curTime.Format(time.Layout))
	fmt.Printf("curTime in Kitchen format=%v\n", curTime.Format(time.Kitchen))
	fmt.Printf("curTime in timestamp format=%v\n", curTime.Format(time.Stamp))
	fmt.Printf("curTime in timestamp with milliseconds format=%v\n", curTime.Format(time.StampMilli))

	// Time arithemtic. Adding date to a time value.
	sameTimeNextYear := curTime.AddDate(1 /* year */, 0, 0)
	fmt.Printf("curTime=%v, sameTimeNextYear=%v\n",
		curTime.Format(time.ANSIC), sameTimeNextYear.Format(time.ANSIC))
}
