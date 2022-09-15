package utils

import (
	"fmt"
	"strings"
)

func Title(s string) {
	line := strings.Repeat("=", 50)
	fmt.Printf("\n%s %s %s\n", line, s, line)
}
