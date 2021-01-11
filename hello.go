package main


import (
	"flag"
	"fmt"
)

const VERSION ="1.0"

func main() {
	version := flag.Bool("v", false, "version du fichier")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
	}

	fmt.Println("Hello, World!\n")
}
