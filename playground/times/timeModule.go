package times

import (
	"fmt"
	"time"
)

func Now() {
	t := time.Now()
	fmt.Println(t)
}
