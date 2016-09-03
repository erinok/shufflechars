// Program shufflechars shuffles the characters of N-copies of a string.
//
// BUG: doesn't handle the unicode combining things right -- e.g., it'll fuck up the red heart emoji.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fatal("usage:", os.Args[0], "NCOPIES STRING")
	}
	rand.Seed(time.Now().Unix())
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fatal(err)
	}
	var runes []rune
	for i := 1; i <= n; i++ {
		runes = append(runes, []rune(os.Args[2])...)
	}
	for i := 0; i < len(runes); i++ {
		k := rand.Intn(len(runes)-i) + i
		runes[i], runes[k] = runes[k], runes[i]
	}
	fmt.Println(string(runes))
}

func fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}
