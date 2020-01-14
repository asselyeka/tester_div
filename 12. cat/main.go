package main

import (
	"io/ioutil"
	"os"
	"io"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	counter := 0
	for range arg {
		counter++
	}
	if counter == 1 {
		for {
			data,_ := ioutil.ReadAll(io.LimitReader(os.Stdin, 1))
			for _, l := range data {
				z01.PrintRune(rune(l))
			}
		}
	} else {

		arguments := []string(arg[1:])

		for _, fileName := range arguments {
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				for _, l := range []rune(err.Error()) {
					z01.PrintRune(l)
				}
				z01.PrintRune('\n')
				return
			}
			for _, l := range string(content) {
				z01.PrintRune(l)
			}
			z01.PrintRune('\n')
		}
	}

}
