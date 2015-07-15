package kace

type link struct {
	val  rune
	link *trie
}

type trie struct {
	nodes []link
}

func newTrie() *trie {
	return &trie{nodes: make([]link, 0)}
}

func (t *trie) add(rs []rune) {
	i := t
	for _, v := range rs {
		ti, ok := searchLink(i.nodes, v)
		if !ok {
			ti = new(trie)
			ti.nodes = make([]link, 0)
			i.nodes = append(i.nodes, link{val: v, link: ti})
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

func searchLink(ls []link, val rune) (*trie, bool) {
	for _, v := range ls {
		if v.val == val {
			return v.link, true
		}
	}
	return nil, false
}
