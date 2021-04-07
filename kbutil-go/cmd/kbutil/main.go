package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cozykeys/resources/kbutil-go/pkg/kbutil"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: kbutil <input-file>\n")
}

func intMain() int {
	if len(os.Args) < 2 {
		printUsage()
		return 1
	}

	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}

	kb := &kbutil.Keyboard{}
	if err := json.Unmarshal(bytes, kb); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}

	svg, err := kb.ToSvg([]string{})
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}
	fmt.Printf("%s\n", svg)

	bytes, err = json.Marshal(kb)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return 1
	}

	fmt.Println(string(bytes))

	return 0
}

func main() {
	os.Exit(intMain())
}
