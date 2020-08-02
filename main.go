package main

import (
	"fmt"
	"os"
	"strings"
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

func main() {
	var err error
	var args = os.Args[1:]
	var p *Parser

	p, err = NewParser()
	if err != nil {
		handleError(err)
	}

	if len(args) == 0 {
		var help []byte
		help, err = Asset("assets/help.txt")
		if err != nil {
			handleError(err)
		}
		fmt.Println(string(help))
		printFile(p, "defs", "Format Attributes")
		printFile(p, "col16", "Colors (16 color palette, prefixes with a lowercase letter)")
		printFile(p, "col256", "Colors (256 color palette, prefixed with a capital letter)")
		return
	}

	err = p.ParseFile(args[0])
	if err != nil {
		handleError(err)
	}

	var res = p.String()
	fmt.Printf("PS1='%s'\n", res)
}

func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func printFile(p *Parser, name, title string) error {
	var data, err = p.loadLines(name)
	if err != nil {
		return err
	}

	fmt.Println(title)
	fmt.Println(strings.Join(data, "\n"))
	fmt.Println("")
	return nil
}
