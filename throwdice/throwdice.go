package main

import (
	"fmt"
	"github.com/tonio-ramirez/dice"
	"os"
	"strconv"
	"strings"
)

func intsToStrings(ints []int) (strings []string) {
	strings = make([]string, len(ints))
	for i, v := range ints {
		strings[i] = strconv.Itoa(v)
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("\t" + os.Args[0] + " <roll description>...")
	} else {
		for i := 1; i < len(os.Args); i++ {
			if roll, err := dice.Roll(os.Args[i]); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%v (%v)\n", roll.Total, strings.Join(intsToStrings(roll.Rolls), ", "))
			}
		}
	}
}
