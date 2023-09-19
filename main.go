package main

import (
	"flag"
	"fmt"
)

func main() {
	path := *flag.String("path", "", "path to basic file")

	flag.Parse()

	if path == "" {
		ePrintln("no path provided, use `-h` to get help")
	} else {
		fmt.Println("path -> ", path)
	}

}

func ePrintln(err string) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	fmt.Println(colorRed + err + colorReset)
}
