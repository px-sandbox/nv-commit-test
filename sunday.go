package main

import (
	"fmt"
	"time"
)

func isWeekday(t time.Time) bool {
	return t.Weekday() > 0 && t.Weekday() < time.Saturday
}

func main() {
	today := time.Now()
	if isWeekday(today) {
		fmt.Println("It's a weekday!")
	} else {
		fmt.Println("It's the weekend!")
	}
}






