package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cozykeys/resources/kbutil-go/pkg/kbutil"
)

func intMain() int {
	fmt.Println("kbutil 2.0!")

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
