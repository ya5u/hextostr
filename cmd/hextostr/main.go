package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ya5u/hextostr"
)

var version string

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	toStr := flag.Bool("toStr", false, "convert hex to ascii")
	toHex := flag.Bool("toHex", false, "convert ascii to hex")
	var length int
	flag.IntVar(&length, "l", 0, "length of hex bytes converted from ascii")
	flag.IntVar(&length, "length", 0, "length of hex bytes converted from ascii")
	flag.Parse()
	args := flag.Args()
	if showVersion {
		fmt.Println("version:", version)
		return
	}
	if !*toStr && *toHex {
		for _, arg := range args {
			h := hextostr.StrToHex(arg, length)
			fmt.Printf("%s -> %s\n", arg, h)
		}
		return
	}
	for _, arg := range args {
		s, err := hextostr.HexToStr(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR on converting %s - %v\n", arg, err)
			continue
		}
		fmt.Printf("%s -> %s\n", arg, s)
	}
}
