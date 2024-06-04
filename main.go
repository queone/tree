package main

import (
	"fmt"
	"os"
	"path"
	"sort"
)

// From https://github.com/kddnewton/tree/tree/main

// The MIT License (MIT)

// Copyright (c) 2016-present Kevin Newton

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// const (
// 	prgname = "tree"
// 	prgver  = "0.0.1"
// )

// // Prints program usage
// func printUsage() {
// 	//X := utl.Red("X")
// 	fmt.Printf(prgname + " v" + prgver + "\n" +
// 		"Directory tree utility - https://github.com/queone/tree\n" +
// 		"Usage: " + prgname + " [options]\n" +
// 		"\n" +
// 		"  -uuid                             Generate new UUID\n" +
// 		"\n" +
// 		"  -id TenantId ClientId Secret      Set up ID for automated login\n" +
// 		"  -tx                               Delete current configured login values and token\n" +
// 		"  -?, -h, --help                    Print this usage page\n")
// 	os.Exit(0)
// }

type Counter struct {
	dirs  int
	files int
}

func (counter *Counter) index(path string) {
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		counter.dirs += 1
	} else {
		counter.files += 1
	}
}

func (counter *Counter) output() string {
	return fmt.Sprintf("\n%d directories, %d files", counter.dirs, counter.files)
}

func dirnamesFrom(base string) []string {
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(err)
	}

	names, _ := file.Readdirnames(0)
	file.Close()

	sort.Strings(names)
	return names
}

func tree(counter *Counter, base string, prefix string) {
	names := dirnamesFrom(base)

	for index, name := range names {
		if name[0] == '.' {
			continue
		}
		subpath := path.Join(base, name)
		counter.index(subpath)

		if index == len(names)-1 {
			fmt.Println(prefix+"└──", name)
			tree(counter, subpath, prefix+"    ")
		} else {
			fmt.Println(prefix+"├──", name)
			tree(counter, subpath, prefix+"│   ")
		}
	}
}

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		directory = "."
	}

	counter := new(Counter)
	fmt.Println(directory)

	tree(counter, directory, "")
	fmt.Println(counter.output())
}
