package inputs

import (
  "bufio"
  "os"
  "strings"
)

func UserInput() []string {
  // this function removes the additional \n character at the end of each
  // line input added by bufio. this also removes whitespaces if any
  text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
  return strings.Split(strings.ReplaceAll(text[:len(text) - 1], " ", ""), ",")
}
