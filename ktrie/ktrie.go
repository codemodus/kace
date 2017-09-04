package ktrie

// KNode ...
type KNode struct {
	val   rune
	end   bool
	links []*KNode
}

// NewKNode ...
func NewKNode(val rune) *KNode {
	return &KNode{
		val:   val,
		links: make([]*KNode, 0),
	}
}

// Add ...
func (n *KNode) Add(rs []rune) {
	cur := n
	for k, v := range rs {
		link := cur.linkByVal(v)
		if link == nil {
			link = NewKNode(v)
			cur.links = append(cur.links, link)
		}

		if k == len(rs)-1 {
			link.end = true
		}

		cur = link
	}
}

// Find ...
func (n *KNode) Find(rs []rune) bool {
	cur := n
	for _, v := range rs {
		cur = cur.linkByVal(v)
		if cur == nil {
			return false
		}
	}

	return cur.end
}

func (n *KNode) linkByVal(val rune) *KNode {
	for _, v := range n.links {
		if v.val == val {
			return v
		}
	}

	return nil
}

// KTrie ...
type KTrie struct {
	maxDepth int
	*KNode
}

// NewKTrie ...
func NewKTrie(data map[string]bool) (*KTrie, error) {
	maxDepth := 0
	n := NewKNode(0)

	for k := range data {
		n.Add([]rune(k))

		if len(k) > maxDepth {
			maxDepth = len(k)
		}
	}

	t := &KTrie{
		maxDepth: maxDepth,
		KNode:    n,
	}

	return t, nil
}

// MaxDepth ...
func (t *KTrie) MaxDepth() int {
	return t.maxDepth
}
