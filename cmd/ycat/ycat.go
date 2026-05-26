package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/goccy/go-yaml"
)

const escape = "\x1b"

func format(attr color.Attribute) string { _ = "STUB: not implemented"; return "" }

func _main(args []string) error { _ = "STUB: not implemented"; return nil }

func main() {
	if err := _main(os.Args); err != nil {
		fmt.Printf("%v\n", yaml.FormatError(err, true, true))
	}
}
