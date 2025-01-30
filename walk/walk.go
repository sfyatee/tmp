//
// Copyright (c) 20XX demian garcia
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	fflag = flag.Bool("f", false, "only non-directories")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	for _, arg := range args {
		err := filepath.WalkDir(arg, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				fmt.Fprintf(os.Stderr, "error accessing %s: %v\n", path, err)
				return nil
			}

			if !*fflag || (*fflag && !d.IsDir()) {
				fmt.Println(path)
			}

			return nil
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "error walking %s: %v\n", arg, err)
		}
	}
}
