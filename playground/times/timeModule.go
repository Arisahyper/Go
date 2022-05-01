package times

import (
	"fmt"
	"time"
)

func Now() {
	t := time.Now()
	fmt.Println(t)
}

func Year(t time.Time) int {
	return t.Year()
}
