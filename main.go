package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/queone/utl"
)

const (
	program_name    = "tree"
	program_version = "0.9.3"
)

func printUsage() {
	n := utl.Yel(program_name)
	v := program_version

	// Reminder: Use tabs for each line in the usage string literal.
	// This allows the following code to replace tabs with spaces and
	// dedent each line appropriately, ensuring consistent indentation.
	// Without using tabs, lines would need to start at column 0, which
	// makes for ugly code.
	usage := fmt.Sprintf(`
	%s v%s
	Directory tree printer - https://github.com/queone/tree
	Usage:
		%s [options] [directory]
	
		Options can be specified in any order. The last specified directory will be used if
		multiple directories are provided.

	Options:
		-f                 Show full file paths. Can be placed before or after the dir path.
		-?, --help, -h     Show this help message and exit

	Examples:
		%s 
		%s -f /path/to/directory
		%s /path/to/directory -f
		%s -h
	`, n, v, n, n, n, n, n)

	// Replace tabs with the specified number of spaces and split into lines
	lines := strings.Split(strings.ReplaceAll(usage, "\t", "  "), "\n")

	// Dedent each line by the specified number of spaces
	for i := range lines {
		if len(lines[i]) >= 2 {
			lines[i] = lines[i][2:]
		}
	}

	fmt.Println(strings.Join(lines, "\n")) // Return the joined back lines
	os.Exit(0)
}

// printTree performs a total of three passes:
// 1) Gather pass: Recursively walks the directory to build a list of entries.
// 2) Determine maximum length: Calculates the longest line (for alignment).
// 3) Print pass: Uses the maximum length to align and print the tree.
func printTree(dir string, showFullPath bool) {
	// entry holds information about each file or directory.
	type entry struct {
		prefix     string // Visual prefix for tree structure
		isLast     bool   // Whether this is the last entry in its level
		name       string // Filename or directory name
		fullPath   string // Absolute path
		isDir      bool   // True if directory
		runeLength int    // Calculated length for alignment
	}

	var all []entry

	// gather recursively walks directories, building a list of entries.
	var gather func(string, string)
	gather = func(curDir, curPrefix string) {
		files, _ := os.ReadDir(curDir)
		var filtered []os.DirEntry
		for _, f := range files {
			if f.Name() != "." && f.Name() != ".." {
				filtered = append(filtered, f)
			}
		}
		for i, f := range filtered {
			isLast := i == len(filtered)-1
			mark := "├── "
			if isLast {
				mark = "└── "
			}
			rawLine := curPrefix + mark + f.Name()
			all = append(all, entry{
				prefix:     curPrefix,
				isLast:     isLast,
				name:       f.Name(),
				fullPath:   filepath.Join(curDir, f.Name()),
				isDir:      f.IsDir(),
				runeLength: utf8.RuneCountInString(rawLine),
			})

			if f.IsDir() {
				nextPrefix := curPrefix + "│   "
				if isLast {
					nextPrefix = curPrefix + "    "
				}
				gather(filepath.Join(curDir, f.Name()), nextPrefix)
			}
		}
	}

	// 1) Gather pass
	gather(dir, "")

	// 2) Determine maximum length
	maxLen := 0
	for _, e := range all {
		if e.runeLength > maxLen {
			maxLen = e.runeLength
		}
	}

	// 3) Print pass
	for _, e := range all {
		mark := "├── "
		if e.isLast {
			mark = "└── "
		}
		coloredName := utl.Gre(e.name)
		if e.isDir {
			coloredName = utl.Blu(e.name)
		}

		line := e.prefix + mark + coloredName
		spacing := (maxLen + 4) - e.runeLength
		if spacing < 1 {
			spacing = 1
		}

		if e.isDir {
			fmt.Println(line)
		} else {
			if showFullPath {
				fmt.Printf("%s%s%s\n", line, strings.Repeat(" ", spacing), utl.Cya(e.fullPath))
			} else {
				fmt.Println(line)
			}
		}
	}
}

func main() {
	showFullPath := false
	var dir string = "."

	args := os.Args[1:]
	if len(args) > 0 {
		// Process command-line arguments in a loop to handle options and directory input.
		// This approach allows for flexible argument ordering, enabling options to be specified
		// before or after the directory. The loop iterates through each argument, setting flags
		// or updating the directory variable as appropriate, ensuring that the last specified
		// directory is used if multiple are provided.
		for _, arg := range args {
			if arg == "-?" || arg == "--help" || arg == "-h" {
				printUsage()
			} else if arg == "-f" {
				showFullPath = true
			} else {
				dir = arg
			}
		}
	}

	printTree(dir, showFullPath)
}
