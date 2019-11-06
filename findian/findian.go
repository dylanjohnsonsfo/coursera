package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}
	toLower := strings.ReplaceAll(strings.ToLower(s), " ", "")

	if strings.Contains(toLower, "i") &&
		strings.Contains(toLower, "a") &&
		strings.Contains(toLower, "n") &&
		strings.HasPrefix(toLower, "i") &&
		strings.HasSuffix(toLower, "n") {
		fmt.Println("Found")

	} else {
		fmt.Println("Not Found")
	}
}
