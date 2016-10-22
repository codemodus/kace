package kace

type node struct {
	val   rune
	end   bool
	links []*node
}

func newNode(val rune, isEnd bool) *node {
	return &node{
		val:   val,
		end:   isEnd,
		links: make([]*node, 0),
	}
}

func (n *node) add(rs []rune) {
	cur := n
	for k, v := range rs {
		isEnd := k == len(rs)-1

		link := cur.linkByVal(v)
		if link == nil {
			link = newNode(v, isEnd)
			cur.links = append(cur.links, link)
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
