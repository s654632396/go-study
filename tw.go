package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	const Padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, Padding, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "bbbb\tccc\tdddd")
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aaa\t123\t4444") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()
}
