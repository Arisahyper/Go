package regex

import (
	"fmt"
	"regexp"
)

func Match() {
	pattern := regexp.MustCompile("^[0-9]{4}$")
	ms := pattern.MatchString("2016")
	fmt.Println(ms)

	match, _ := regexp.MatchString("^[0-9]{4}$", "1234")
	fmt.Println(match)
}
