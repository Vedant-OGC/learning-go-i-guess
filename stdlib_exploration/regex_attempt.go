package stdlibexploration

import (
	"fmt"
	"regexp"
)

func RegexAttempt() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("find string:", r.FindString("peach punch"))
	fmt.Println("submatch:", r.FindStringSubmatch("peach")) // [peach ea]
}
