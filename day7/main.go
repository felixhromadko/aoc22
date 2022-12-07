package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type File struct {
	name   string
	parent *File

	contentSize int
	children    []*File
}

func (f *File) Size() int {
	if len(f.children) != 0 {
		var size int
		for _, child := range f.children {
			size += child.Size()
		}
		return size
	}

	return f.contentSize
}

func (f *File) RecursiveLs(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("- %s: (%d)\n", f.name, f.Size())

	for _, child := range f.children {
		child.RecursiveLs(level + 1)
	}
}

func SumOfSmallFolders(f *File) int {
	var sum int
	if len(f.children) > 0 && f.Size() < 100000 {
		sum += f.Size()
	}

	for _, child := range f.children {
		sum += SumOfSmallFolders(child)
	}

	return sum
}

func FindSmallestFolderSizeBiggerThan(biggerThan int, f *File) int {
	bestMatch := 999999999999999999

	if len(f.children) > 0 {
		size := f.Size()
		if size >= biggerThan && bestMatch > size {
			bestMatch = size
		}
	}

	for _, child := range f.children {
		bestSubMatch := FindSmallestFolderSizeBiggerThan(biggerThan, child)
		if bestSubMatch < bestMatch {
			bestMatch = bestSubMatch
		}
	}

	return bestMatch
}

func main() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	lineRegex := regexp.MustCompile("\\$ ([a-z]+) ?(.+)?")
	lines := strings.Split(string(inputRaw), "\n")

	rootFolder := &File{
		name: "root",
	}
	currentFolder := rootFolder
lineFor:
	for i := 0; i < len(lines); i++ {
		// if starts with $, it's a command
		if !strings.HasPrefix(lines[i], "$") {
			panic("invalid line")
		}

		// parse the line
		submatch := lineRegex.FindStringSubmatch(lines[i])

		switch submatch[1] {
		case "cd":
			if strings.HasPrefix(submatch[2], "/") {
				// treat as absolute path
				currentFolder = rootFolder
				continue
			}

			if submatch[2] == ".." {
				currentFolder = currentFolder.parent
				continue
			}

			for _, child := range currentFolder.children {
				if child.name == submatch[2] {
					currentFolder = child
					continue lineFor
				}
			}
			panic("folder not found")
		case "ls":
			for {
				i++
				if i >= len(lines) {
					break
				}

				if strings.HasPrefix(lines[i], "$") {
					i--
					break
				}
				parts := strings.Split(lines[i], " ")

				if parts[0] == "dir" {
					currentFolder.children = append(currentFolder.children, &File{
						name:   parts[1],
						parent: currentFolder,
					})
				} else {
					fileSize, err := strconv.Atoi(parts[0])
					checkErr(err)

					currentFolder.children = append(currentFolder.children, &File{
						name:        parts[1],
						contentSize: fileSize,
					})
				}
			}
		}
	}

	rootFolder.RecursiveLs(0)

	fmt.Println(SumOfSmallFolders(rootFolder))

	totalSpace := 70000000
	spaceNeeded := 30000000
	spaceFree := totalSpace - rootFolder.Size()

	fmt.Println("Need ", spaceNeeded-spaceFree)

	// find the smallest possible folder to delete
	fmt.Println(FindSmallestFolderSizeBiggerThan(spaceNeeded-spaceFree, rootFolder))

}
