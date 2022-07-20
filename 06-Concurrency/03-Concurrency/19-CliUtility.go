//21. Cli Utility Two

//good utilities
//http://patorjk.com/software/taag/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func showBanner() {
	auther := "mrAlishah"
	version := "1.0"
	name := fmt.Sprintf("CLI utility (v.%s)", version)

	banner := `
 ________       .__                         
 /  _____/  ____ |  | _____    ____    ____  
/   \  ___ /  _ \|  | \__  \  /    \  / ___\ 
\    \_\  (  <_> )  |__/ __ \|   |  \/ /_/  >
 \______  /\____/|____(____  /___|  /\___  / 
        \/                 \/     \//_____/  
 `

	fmt.Println(banner)
	all_lines := strings.Split(banner, "\n")
	len_line := len(all_lines[4])

	color.Green(fmt.Sprintf("%[1]*s", (len_line+len(name))/2, name))
	color.Blue(fmt.Sprintf("%[1]*s", (len_line+len(auther))/2, auther))
}

//guide for formatting
//https://www.cs.ubc.ca/~bestchai/teaching/cs416_2015w2/go1.4.3-docs/pkg/fmt/index.html
// color.Green(fmt.Sprintf("%[2]d %[1]d\n", 11, 22))
//	color.Green(fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 3, 6))
func main() {
	showBanner()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(">>>")
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Println(cmd)

		switch cmd {
		case "something":
			fmt.Println("command something is execute")
			break
		}

		fmt.Print(">>>")

	}

	if scanner.Err() != nil {
		//Handle error
	}
}
