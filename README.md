# tree

A lightweight directory tree printing utility.

## Why?

Why another tree utility??

**Availability**: If the official tree utility is not installed or hard to get, you can quickly compile this with Go.
**Cross-Platform**: Easily compile on Linux, Mac, or Windows.
**Simplicity**: Covers 99% of typical use cases with a minimal feature set.
**Learning Opportunity**: A great way to practice coding in Go.


## Getting Started

To compile, clone this repository and run:
```bash
./build_go
```

This version maintains the essential information while being more concise and easier to read for advanced users.


## Usage

```bash
$ tree -?

tree v0.9.1
Directory tree printer - https://github.com/queone/tree
Usage:
  tree [options] [directory]

  Options can be specified in any order. The last specified directory will be used if multiple directories are provided.

Options:
  -f                 Show full file paths. Can be placed before or after the directory path.
  -?, --help, -h     Show this help message and exit

Examples:
  tree
  tree -f /path/to/directory
  tree /path/to/directory -f
  tree -h
```
