package stdlibexploration

import (
	"fmt"
	"regexp"
)

func RegexAttempt() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("match peach:", match)
}
