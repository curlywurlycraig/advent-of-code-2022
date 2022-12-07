package day7

import (
	"strconv"
	"strings"
)

const (
	TypeDir = iota
	TypeFile
)

type FsNode struct {
	Type int
	Size int
	Name string

	Parent   *FsNode
	Contents map[string]*FsNode
}

func Day7(input string) int {
	rootNode := buildFsTree(input)

	// Search tree for dirs less than magic number 100000
	result := 0
	searchQueue := []*FsNode{rootNode}
	currNode := rootNode
	for len(searchQueue) > 0 {
		// Pop from queue
		currNode, searchQueue = searchQueue[0], searchQueue[1:]
		if currNode.Size < 100_000 {
			result += currNode.Size
		}

		for _, n := range currNode.Contents {
			if n.Type == TypeDir {
				searchQueue = append(searchQueue, n)
			}
		}
	}

	return result
}

func Day7Part2(input string) int {
	rootNode := buildFsTree(input)

	diskSize := 70_000_000
	requiredSpace := 30_000_000

	currentUnused := diskSize - rootNode.Size
	requiredToDelete := requiredSpace - currentUnused

	// Search tree for dirs less than magic number 100000
	minSatisfactory := rootNode.Size
	searchQueue := []*FsNode{rootNode}
	currNode := rootNode
	for len(searchQueue) > 0 {
		// Pop from queue
		currNode, searchQueue = searchQueue[0], searchQueue[1:]
		if currNode.Size < minSatisfactory && currNode.Size > requiredToDelete {
			minSatisfactory = currNode.Size
		}

		for _, n := range currNode.Contents {
			if n.Type == TypeDir {
				searchQueue = append(searchQueue, n)
			}
		}
	}

	return minSatisfactory
}

func buildFsTree(input string) *FsNode {
	rootNode := &FsNode{
		Type:     TypeDir,
		Size:     -1,
		Name:     "/",
		Parent:   nil,
		Contents: make(map[string]*FsNode),
	}

	currNode := rootNode

	lines := strings.Split(input, "\n")
	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		line := lines[lineIdx]

		if strings.HasPrefix(line, "$ cd ") {
			navigationDirName := line[len("$ cd "):]
			if navigationDirName == "/" {
				continue
			}

			if navigationDirName == ".." {
				currNode = currNode.Parent
				continue
			}

			currNode = currNode.Contents[navigationDirName]
		} else if strings.HasPrefix(line, "$ ls") {
			// Noop
		} else {
			newNode := parseLsLine(line, currNode)
			currNode.Contents[newNode.Name] = newNode
		}
	}

	rootNode.CalculateSizes()
	return rootNode
}

func parseLsLine(line string, currDir *FsNode) *FsNode {
	if strings.HasPrefix(line, "dir") {
		name := line[4:]
		return &FsNode{
			Type:     TypeDir,
			Size:     -1,
			Name:     name,
			Parent:   currDir,
			Contents: make(map[string]*FsNode),
		}
	} else {
		parts := strings.Split(line, " ")
		size, _ := strconv.Atoi(parts[0])
		name := parts[1]

		return &FsNode{
			Type:     TypeFile,
			Size:     size,
			Name:     name,
			Parent:   currDir,
			Contents: make(map[string]*FsNode),
		}
	}
}

func (n *FsNode) CalculateSizes() {
	if n.Size != -1 {
		return
	}

	result := 0
	for _, c := range n.Contents {
		c.CalculateSizes()
		result += c.Size
	}
	n.Size = result
}
