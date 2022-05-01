package regex

import (
	"fmt"
	"regexp"
)

func Match() {
	match, _ := regexp.MatchString("^[0-9]{4}$", "1234")
	fmt.Println(match)
}
