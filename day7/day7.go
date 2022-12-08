package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

const (
	CommandIdentifier = "$"
	ChangeDirectory   = "cd"
	ListDirectory     = "ls"
	RootNodeName      = "/"
	MoveUpIdentifier  = ".."
)

const (
	TotalDiskSpace      = 70000000
	UpdateSpaceRequired = 30000000
)

func Solution(sessionCookie, pt1Text, pt2Text string) {
	rows := util.GetRows("https://adventofcode.com/2022/day/7/input", sessionCookie)
	rootNode := parseCommandOutputIntoTreeStructure(rows)

	fmt.Printf(pt1Text)
	fmt.Printf("%d\n", findDirsWithSizeUpTo(rootNode, 100000))

	fmt.Printf(pt2Text)
	fmt.Printf("Used space: %d\tspace to free up: %d\n", rootNode.GetSize(), rootNode.GetSize()-(TotalDiskSpace-UpdateSpaceRequired))

	directories := findSmallestDirToDelete(rootNode, rootNode.GetSize()-(TotalDiskSpace-UpdateSpaceRequired))
	sort.Sort(util.Dirs(directories))
	fmt.Printf("Smallest dir to delete: \\%s (%d)\n", directories[0].Name, directories[0].GetSize())
}

func parseCommandOutputIntoTreeStructure(input []string) util.Node {
	rootNode := &util.DirNode{Name: RootNodeName, Children: []util.Node{}}
	currentNode := rootNode
	for _, row := range input {
		if isChangeDirectoryCommand(row) {
			to := strings.Split(row, " ")[2]
			currentNode = changeDirectory(rootNode, currentNode, to)
			if currentNode.Children == nil {
				currentNode.Children = []util.Node{}
			}
		} else if isListCommand(row) {
			// Do nothing.
		} else {
			// Update children of current node.
			child := getChildNode(row)
			child.SetParent(currentNode)
			currentNode.Children = append(currentNode.Children, child)
			incrementDirectorySize(currentNode, child.GetSize())
		}
	}
	return rootNode
}

func isChangeDirectoryCommand(input string) bool {
	parts := strings.Split(input, " ")
	return parts[0] == CommandIdentifier && parts[1] == ChangeDirectory
}

func isListCommand(input string) bool {
	parts := strings.Split(input, " ")
	return parts[0] == CommandIdentifier && parts[1] == ListDirectory
}

func changeDirectory(rootNode, currentNode *util.DirNode, changeTo string) *util.DirNode {
	if changeTo == RootNodeName {
		return rootNode
	} else if changeTo == MoveUpIdentifier {
		if currentNode.GetParent() == rootNode {
			return rootNode
		} else {
			return currentNode.GetParent().(*util.DirNode)
		}
	} else {
		node := getDirectoryNode(currentNode, changeTo)
		if node == nil {
			fmt.Printf("Nil node, no change dir: '%s' found\n", changeTo)
			return nil
		} else {
			return node
		}
	}
}

func getDirectoryNode(currentNode util.Node, changeTo string) *util.DirNode {
	if !currentNode.IsDirectory() {
		fmt.Println("currentNode isn't a directory!")
		return nil
	}
	for _, node := range currentNode.GetChildren() {
		if node.IsDirectory() && node.GetName() == changeTo {
			return node.(*util.DirNode)
		}
	}
	fmt.Println("nothing found...")
	return nil
}

func getChildNode(input string) util.Node {
	parts := strings.Split(input, " ")
	if parts[0] == "dir" {
		return &util.DirNode{Name: parts[1]}
	} else {
		size, _ := strconv.Atoi(parts[0])
		return &util.FileNode{Name: parts[1], Size: size}
	}
}

func incrementDirectorySize(node *util.DirNode, size int) {
	node.SetSize(node.GetSize() + size)
	if node.GetParent() != nil {
		incrementDirectorySize(node.GetParent().(*util.DirNode), size)
	}
}

func findDirsWithSizeUpTo(node util.Node, inclusiveUpperBound int) int {
	result := 0
	for _, child := range node.GetChildren() {
		if !child.IsDirectory() {
			continue
		}
		if child.GetSize() <= inclusiveUpperBound {
			result += child.GetSize()
		}
		result += findDirsWithSizeUpTo(child, inclusiveUpperBound)
	}
	return result
}

func findSmallestDirToDelete(node util.Node, spaceToFree int) []*util.DirNode {
	result := []*util.DirNode{}
	if node.IsDirectory() && node.GetSize() >= spaceToFree {
		result = append(result, node.(*util.DirNode))
	}
	for _, child := range node.GetChildren() {
		if child.IsDirectory() {
			result = append(result, findSmallestDirToDelete(child, spaceToFree)...)
		}
	}
	return result
}
