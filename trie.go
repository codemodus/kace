package kace

type node struct {
	val  rune
	link *trie
}

type trie struct {
	nodes []*node
}

func newTrie() *trie {
	return &trie{nodes: make([]*node, 0)}
}

func (t *trie) add(rs []rune) {
	i := t
	for _, v := range rs {
		ti, ok := searchLink(i.nodes, v)
		if !ok {
			ti = newTrie()
			i.nodes = append(i.nodes, &node{val: v, link: ti})
		}
		i = ti
	}
}

func (t *trie) find(rs []rune) bool {
	i := t
	for _, v := range rs {
		ti, ok := searchLink(i.nodes, v)
		if !ok {
			return false
		}
		i = ti
	}
	return true
}

func searchLink(ls []*node, val rune) (*trie, bool) {
	for _, v := range ls {
		if v.val == val {
			return v.link, true
		}
	}
	return nil, false
}
