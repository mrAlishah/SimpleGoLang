//20. Cli Utility One
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//use io.Reader as a string type
//<-chan read only string channel as a return type
func read(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		s := bufio.NewScanner(r)
		for s.Scan() {
			lines <- s.Text()
		}
	}()

	return lines
}

func main() {
	mes := read(os.Stdin)

	for anu := range mes {
		fmt.Println("Msg out: ", anu)

		switch anu {
		case "hello":
			fmt.Println("world!")
			break
		default:
			break
		}
	}
}
