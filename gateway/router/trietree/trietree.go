package trietree

type Node struct {
	value interface{}
	child map[string]*Node
}

type TrieTree struct {
	root *Node
}

func NewTrieTree() *TrieTree {
	return &TrieTree{root: &Node{child: make(map[string]*Node)}}
}

func (t *TrieTree) Add(path []string, value interface{}) {
	var node = t.root
	if len(path) == 0 {
		node.value = value
		return
	}
	for i := 0; i < len(path); i++ {
		subPath := path[i]
		if subPath == "*" {
			newNode := &Node{value: value, child: make(map[string]*Node)}
			node.child["*"] = newNode
			node = newNode
		} else {
			child := node.child[subPath]
			if child == nil {
				newNode := &Node{child: make(map[string]*Node)}
				node.child[subPath] = newNode
				node = newNode
			} else {
				node = child
			}
		}
		if i == len(path)-1 {
			node.value = value
		}
	}
}

func (t *TrieTree) Del(path []string) interface{} {
	var ret interface{} = nil
	node := t.root
	if len(path) == 0 {
		ret := node.value
		node.value = nil
		return ret
	}
	for i := 0; i < len(path); i++ {
		subPath := path[i]
		if subPath == "*" {
			ans := node.child["*"]
			delete(node.child, "*")
			return ans.value
		} else {
			child := node.child[subPath]
			if child == nil {
				return nil
			}
			node = child
		}
		if i == len(path)-1 {
			ret = node.value
			node.value = nil
		}
	}
	return ret
}

func (t *TrieTree) Get(path []string) interface{} {
	var ret interface{} = nil
	var node = t.root
	if node.child["*"] != nil {
		ret = node.child["*"].value
	}
	if len(path) == 0 && node.value != nil {
		ret = node.value
	}
	for i := 0; i < len(path); i++ {
		subPath := path[i]
		if node.child["*"] != nil {
			ret = node.child["*"].value
		}
		if node.child[subPath] == nil {
			return ret
		} else {
			node = node.child[subPath]
		}
		if i == len(path)-1 && node.value != nil {
			ret = node.value
		}
	}
	return ret
}
