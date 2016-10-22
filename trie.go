package kace

type node struct {
	val   rune
	end   bool
	links []*node
}

func newNode(val rune) *node {
	return &node{
		val:   val,
		links: make([]*node, 0),
	}
}

func (n *node) add(rs []rune) {
	cur := n
	for k, v := range rs {

		link := cur.linkByVal(v)
		if link == nil {
			link = newNode(v)
			cur.links = append(cur.links, link)
		}

		if k == len(rs)-1 {
			link.end = true
		}

		cur = link
	}
}

func (n *node) find(rs []rune) bool {
	cur := n
	for _, v := range rs {
		cur = cur.linkByVal(v)
		if cur == nil {
			return false
		}
	}

	return cur.end
}

func (n *node) linkByVal(val rune) *node {
	for _, v := range n.links {
		if v.val == val {
			return v
		}
	}

	return nil
}

type trie struct {
	maxDepth int
	*node
}

func newTrie(data map[string]bool) *trie {
	maxDepth := 0
	n := newNode(0)
	for k := range data {
		n.add([]rune(k))

		if len(k) > maxDepth {
			maxDepth = len(k)
		}
	}

	return &trie{maxDepth: maxDepth, node: n}
}
