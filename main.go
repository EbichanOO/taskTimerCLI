package main

import (
	"fmt"
	"time"
)

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
	return marks[i%4]
}

func getNowTime() time.Time {
	return time.Now()
}

func main() {
	cnt := 5000000000
	for i := 1; i <= cnt; i++ {
		if i%(cnt/100) == 0 {
			p := i / (cnt / 100)
			fmt.Printf("\rLoading: %s %2d%%", mark(p), p)
		}
	}
	fmt.Println("\nDone.")
}
