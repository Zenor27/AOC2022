package main

import (
	"aoc2022/utils"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	utils.Run("day07", functionPart1, functionPart2)
}

type Folder struct {
	name         string
	parentFolder *Folder
	files        []File
	folders      []Folder
}

type File struct {
	name string
	size int
}

func (folder Folder) sumSizes() int {
	size := 0
	for _, file := range folder.files {
		size += file.size
	}
	for _, subFolder := range folder.folders {
		size += subFolder.sumSizes()
	}
	return size
}

func (folder Folder) computeAllSizes() []int {
	total := folder.sumSizes()
	sizes := make([]int, 0)
	sizes = append(sizes, total)
	for _, subFolder := range folder.folders {
		sizes = append(sizes, subFolder.computeAllSizes()...)
	}
	return sizes
}

const maxFolderSize = 100000

func getRootFolder(input string) Folder {
	lines := strings.Split(input, "\n")
	lines = lines[1:]

	rootFolder := Folder{name: "/", parentFolder: nil}
	currentFolder := &rootFolder

	idx := 0
	for idx < len(lines) {
		line := lines[idx]
		if strings.Contains(line, "cd") {
			newFolder := strings.Split(line, " ")[2]
			if newFolder == ".." {
				currentFolder = currentFolder.parentFolder
			} else {
				// Find the folder in the current folder
				folderIdx := slices.IndexFunc(currentFolder.folders, func(f Folder) bool {
					return f.name == newFolder
				})
				currentFolder = &currentFolder.folders[folderIdx]
			}
		} else if strings.Contains(line, "ls") {
			for idx+1 < len(lines) && !strings.Contains(lines[idx+1], "$") {
				idx++
				info := lines[idx]
				// It is a file
				if !strings.Contains(info, "dir") {
					file := strings.Split(info, " ")
					size, _ := strconv.Atoi(file[0])
					name := file[1]
					currentFolder.files = append(currentFolder.files, File{name: name, size: size})

				} else {
					// It is a folder
					name := strings.Split(info, " ")[1]
					folder := Folder{name: name, parentFolder: currentFolder}
					currentFolder.folders = append(currentFolder.folders, folder)
				}
			}
		}

		idx++
	}

	return rootFolder
}

func functionPart1(input string) int {
	totalSizes := getRootFolder(input).computeAllSizes()

	total := 0
	for _, size := range totalSizes {
		if size <= maxFolderSize {
			total += size
		}
	}
	return total
}

const totalDiskSpace = 70000000
const requiredToUpdate = 30000000

func functionPart2(input string) int {
	rootFolder := getRootFolder(input)
	sizes := rootFolder.computeAllSizes()
	maxSize := sizes[0]
	availableSpace := totalDiskSpace - maxSize

	sort.Ints(sizes)
	for _, size := range sizes {
		claimable := size + availableSpace
		if claimable > requiredToUpdate {
			return size
		}
	}

	return maxSize

}
