// Program shufflechars shuffles the characters of N-copies of a string.
//
// BUG: just smart enough to handle emoji; likely breaks on fancier unicode stuff (isn't working on the grapheme level like it should).
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode"
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
	// shuffle
	for i := 0; i < len(runes); {
		k := rand.Intn(len(runes)-i) + i
		// we need to move modifier characters as a contiguous unit. we only handle a single modifier correctly.
		if ismod(runes[k]) && k-1 >= i {
			runes[i], runes[k-1] = runes[k-1], runes[i]
			runes[i+1], runes[k] = runes[k], runes[i+1]
			i += 2
		} else if k+1 < len(runes) && ismod(runes[k+1]) {
			runes[i], runes[k] = runes[k], runes[i]
			runes[i+1], runes[k+1] = runes[k+1], runes[i+1]
			i += 2
		} else {
			runes[i], runes[k] = runes[k], runes[i]
			i++
		}
	}
	fmt.Println(string(runes))
	// fmt.Println([]rune(runes))
}

var mods = []*unicode.RangeTable{unicode.Variation_Selector, unicode.Sk}

func ismod(r rune) bool {
	return unicode.IsOneOf(mods, r)
}

func fatal(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}
