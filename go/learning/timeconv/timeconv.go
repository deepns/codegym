// Exploring time parsing and formatting
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Get current time
	cTime := time.Now()

	// Default reference format "2006-01-02 15:04:05.999999999 -0700 MST"
	fmt.Printf("cTime: %v\n", cTime)

	// Constants in time has additional predefined formatting strings.
	fmt.Printf("cTime.Format(time.ANSIC): %v\n", cTime.Format(time.ANSIC))
	fmt.Printf("cTime.Format(time.RFC1123): %v\n", cTime.Format(time.RFC1123))
	fmt.Printf("cTime.Format(time.Layout): %v\n", cTime.Format(time.Layout))
	fmt.Printf("cTime.Format(time.Kitchen): %v\n", cTime.Format(time.Kitchen))
	fmt.Printf("cTime.Format(time.Stamp): %v\n", cTime.Format(time.Stamp))

	// Using custom format
	// formatting is based on reference time
	fmt.Printf("cTime.Format(\"Mon, Jan 02 2006\"): %v\n", cTime.Format("Mon, Jan 02 2006"))
	fmt.Printf("cTime.Format(\"Jan _2, 06\"): %v\n", cTime.Format("Jan _2, 06"))

	aTime := cTime.AddDate(0, 0, -50)
	fmt.Printf("aTime.Format(\"_1/_2 2006\"): %v\n", aTime.Format("01/2 2006"))

	t1, err := time.Parse("01/02 2006", "02/26 2022")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("t1: %v\n", t1)

	t1, err = time.Parse("Jan 01 2006 03:04PM", "Oct 10 2021 05:04PM")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("t1: %v\n", t1)

	// If no date is specified, it defaults to 0000-01-01
	kitchenTime, _ := time.Parse(time.Kitchen, "12:11PM")
	fmt.Printf("kitchenTime: %v\n", kitchenTime)

	// Convert unix epoch seconds to timestamp
	billionthSecond := time.Unix(1e9, 0)
	fmt.Printf("billionthSecond: %v\n", billionthSecond)
	fmt.Printf("billionthSecond.Unix(): %v\n", billionthSecond.Unix())
}
