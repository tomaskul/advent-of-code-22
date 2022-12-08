package util

type Node interface {
	GetName() string
	IsDirectory() bool
	GetChildren() []Node
	SetParent(Node)
	GetParent() Node
	SetSize(int)
	GetSize() int
}

type DirNode struct {
	Name         string
	Children     []Node
	parent       Node
	ChildrenSize int
}

func (n *DirNode) GetName() string {
	return n.Name
}

func (n *DirNode) IsDirectory() bool {
	return true
}

func (n *DirNode) GetChildren() []Node {
	return n.Children
}

func (n *DirNode) SetParent(parent Node) {
	n.parent = parent
}

func (n *DirNode) GetParent() Node {
	return n.parent
}

func (n *DirNode) SetSize(size int) {
	n.ChildrenSize = size
}

func (n *DirNode) GetSize() int {
	return n.ChildrenSize
}

type Dirs []*DirNode

// implement the functions from the sort.Interface
func (d Dirs) Len() int {
	return len(d)
}

func (d Dirs) Less(i, j int) bool {
	return d[i].GetSize() < d[j].GetSize()
}

func (d Dirs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type FileNode struct {
	Name   string
	Size   int
	parent Node
}

func (n *FileNode) GetName() string {
	return n.Name
}

func (n *FileNode) IsDirectory() bool {
	return false
}

func (n *FileNode) GetChildren() []Node {
	return nil
}

func (n *FileNode) SetParent(parent Node) {
	n.parent = parent
}

func (n *FileNode) GetParent() Node {
	return n.parent
}

func (n *FileNode) SetSize(size int) {
	n.Size = size
}

func (n *FileNode) GetSize() int {
	return n.Size
}
