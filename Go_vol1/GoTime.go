package main

import (
	"fmt"
	"time"
)

func main() {
	present := time.Now()

	fmt.Println(present)
	DOB := time.Date(1993, 02, 28, 9, 04, 39, 213, time.Local)
	fmt.Println(DOB)
	fmt.Println(DOB.Year())
	fmt.Println(DOB.Month())
	fmt.Println(DOB.Day())
	fmt.Println(DOB.Hour())
	fmt.Println(DOB.Minute())
	fmt.Println(DOB.Second())
	fmt.Println(DOB.Nanosecond())
	fmt.Println(DOB.Location())

	fmt.Println(DOB.Weekday())

	fmt.Println(DOB.Before(present))
	fmt.Println(DOB.After(present))
	fmt.Println(DOB.Equal(present))

	diff := present.Sub(DOB)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())

	fmt.Println(DOB.Add(diff))

	fmt.Println(present.Format("Mon, Jan 2, 2006 at 3:04pm"))

	current_time := time.Now()
	fmt.Println(current_time)
	sec := current_time.Unix()
	nanos := current_time.UnixNano()

	fmt.Println(sec)
	fmt.Println(nanos)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nanos))

	for _ = range time.Tick(1 * time.Second) {
		fmt.Println("\t B : Okay!")
	}

}
