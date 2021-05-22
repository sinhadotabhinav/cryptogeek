package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UserInput func removes the additional \n character at the end of each line input
// added by bufio. this also removes whitespaces if any
func UserInput() []string {
	fmt.Print("> ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Split(strings.ReplaceAll(text[:len(text)-1], " ", ""), ",")
}
