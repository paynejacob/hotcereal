package graph

type Graph struct {
	root *Node
}

func New() *Graph {
	return &Graph{
		root: newNode(0),
	}
}

func (g *Graph) Write(tokens []string, value string) {
	for i := 0; i < len(tokens); i++ {
		g.writeToken(tokens[i], value)
	}
}

func (g *Graph) Search(q string) []string {
	var root *Node

	root = g.root
	for _, char := range q {
		// check if this sequence value exists
		if root.HasChild(char) {
			root = root.children[char]
		} else {
			return []string{}
		}
	}

	return root.Values()
}

func (g *Graph) Delete(value string) {
	var stack nodeStack
	var root *Node

	stack.Push(g.root)
	for {

		// We have traversed all relevant nodes exit
		if stack.Empty() {
			break
		}

		// get the next node
		root = stack.Pop()

		// if this node has the index it's children may also so we need to check them
		if root.HasValue(value) {
			for k := range root.children {
				stack.Push(root.children[k])
			}
		} else {
			// if this node does not have the value we are done
			continue
		}

		root.RemoveValue(value)

		if root.Empty() {
			if root.parent != nil {
				root.parent.RemoveChild(root.char)
			}
		}
	}
}

func (g *Graph) writeToken(token string, value string) {
	var root *Node
	var char rune

	root = g.root
	root.AddValue(value)
	for _, char = range token {
		// ensure this Node exists
		if !root.HasChild(char) {
			root.AddChild(newNode(char))
		}

		// move down the tree
		root = root.children[char]

		// write value to Node
		root.AddValue(value)
	}
}
