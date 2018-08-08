package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/go-logfmt/logfmt"
	flag "github.com/spf13/pflag"
)

func main() {

	flag.BoolVar(&color.NoColor, "no-color", false, "No Colorized Output")

	flag.ErrHelp = errors.New("")
	flag.Parse()

	fields := flag.Args()
	d := logfmt.NewDecoder(bufio.NewReader(os.Stdin))

	var b bytes.Buffer
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 4, ' ', 0)

	for d.ScanRecord() {
		line := make(map[string]string)

		for d.ScanKeyval() {
			line[string(d.Key())] = string(d.Value())
		}

		for _, field := range fields {
			v, exists := line[field]
			if exists {
				b.WriteString(color.BlueString(field) + "=")
				if strings.Contains(v, " ") {
					v = "\"" + v + "\""
				}
				b.WriteString(v + "\t")
			}
		}
		b.WriteString("\n")

		w.Write(b.Bytes())
		b.Reset()
	}

	w.Flush()
}
