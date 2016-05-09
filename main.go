// psgrep
//
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var psOpts string = "-ef"
	var grep, header string
	var grepFound []string
	var out bytes.Buffer

	if os.Getenv("PSGREP_OPTS") != "" {
		psOpts = os.Getenv("PSGREP_OPTS")
	}

	if len(os.Args) > 1 {
		grep = strings.Join(os.Args[1:], " ")
	}

	cmd := exec.Command("ps", psOpts)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range strings.Split(out.String(), "\n") {
		if i == 0 {
			header = s
		}
		if ok := strings.Contains(s, "psgrep"); ok {
			continue
		}
		if ok := strings.Contains(s, grep); ok {
			grepFound = append(grepFound, s)
		}
	}
	if len(grepFound) > 0 {
		fmt.Println(header)
		for _, s := range grepFound {
			fmt.Println(s)
		}
	} else {
		os.Exit(1)
	}
}
