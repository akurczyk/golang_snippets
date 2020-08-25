package main

import (
	"fmt"
	"time"
)

func main() {
	// Build
	now := time.Now()
	fmt.Println(now)

	someTimeAgo := time.Date(1994, 12, 10, 22, 30, 45, 500000, time.UTC)
	fmt.Println(someTimeAgo)

	// Extract
	fmt.Println("Year:", someTimeAgo.Year())
	fmt.Println("Month:", someTimeAgo.Month())
	fmt.Println("Day:", someTimeAgo.Day())
	fmt.Println("Hour:", someTimeAgo.Hour())
	fmt.Println("Minute:", someTimeAgo.Minute())
	fmt.Println("Second:", someTimeAgo.Second())
	fmt.Println("Nanosecond:", someTimeAgo.Nanosecond())
	fmt.Println("YearDay:", someTimeAgo.YearDay())
	fmt.Println("Weekday:", someTimeAgo.Weekday())

	year, month, day := someTimeAgo.Date()
	fmt.Println("Date: ", year, month, day)

	hour, minute, second := someTimeAgo.Clock()
	fmt.Println("Clock: ", hour, minute, second)

	json, _ := someTimeAgo.MarshalJSON()
	fmt.Println("JSON: ", string(json))

	// Compare
	fmt.Println(someTimeAgo.Before(now))
	fmt.Println(someTimeAgo.After(now))
	fmt.Println(someTimeAgo.Equal(now))

	diff := now.Sub(someTimeAgo)
	fmt.Println(diff)

	fmt.Println(someTimeAgo.Add(diff))
	fmt.Println(someTimeAgo.Add(diff * 2))
}
