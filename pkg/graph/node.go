package graph

type Node struct {
	char rune

	parent   *Node
	children map[rune]*Node
	values   map[string]bool
}

func newNode(char rune) *Node {
	return &Node{char: char, children: make(map[rune]*Node, 0), values: make(map[string]bool, 0)}
}

func (n *Node) AddValue(value string) {
	n.values[value] = true
}

func (n *Node) HasValue(value string) bool {
	var found bool

	_, found = n.values[value]

	return found
}

func (n *Node) RemoveValue(value string) {
	delete(n.values, value)
}

func (n *Node) Values() []string {
	rvalues := make([]string, len(n.values))

	i := 0
	for k, _ := range n.values {
		rvalues[i] = k
		i++
	}

	return rvalues
}

func (n *Node) Empty() bool {
	return len(n.values) == 0
}

func (n *Node) AddChild(child *Node) {
	child.parent = n
	n.children[child.char] = child
}

func (n *Node) HasChild(char rune) bool {
	if _, ok := n.children[char]; ok {
		return true
	}

	return false
}

func (n *Node) RemoveChild(char rune) {
	delete(n.children, char)
}
