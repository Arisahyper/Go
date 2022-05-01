package regex

import (
	"fmt"
	"regexp"
)

func Match() {
	pattern := regexp.MustCompile(`^0\d{9,10}$`)
	ms := pattern.MatchString("090-1234-5678")
	ms2 := pattern.MatchString("09012345678")
	fmt.Println(ms)
	fmt.Println(ms2)

	match, _ := regexp.MatchString("^[0-9]{4}$", "1234")
	fmt.Println(match)
}
