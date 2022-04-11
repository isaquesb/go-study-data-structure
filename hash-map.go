package main

type HashMap struct {
	index  map[int64]*HashNode
	hashFn func(string) int64
}

type HashNode struct {
	node Node
	key  string
}

type HashConfig struct {
	Fn   func(string) int64
	Size int
}

func NewHashMap(cfg HashConfig) *HashMap {
	if nil == cfg.Fn {
		cfg.Fn = Djb2
	}
	return &HashMap{
		index:  make(map[int64]*HashNode, cfg.Size),
		hashFn: cfg.Fn,
	}
}

func (h *HashMap) Has(key string) bool {
	hashed := h.hashedKey(key)
	has, _ := h.getNode(hashed)
	return has
}

func (h *HashMap) Put(key string, value NodeContent) {
	hashed := h.hashedKey(key)
	node := &HashNode{
		key:  key,
		node: Node{Data: value},
	}
	h.index[hashed] = node
}

func (h *HashMap) Delete(key string) (bool, NodeContent) {
	hashed := h.hashedKey(key)
	has, hashNode := h.getNode(hashed)
	if !has {
		return false, nil
	}
	h.index[hashed] = nil
	return true, hashNode.node.Data
}

func (h *HashMap) Search(key string) (bool, NodeContent) {
	hashed := h.hashedKey(key)
	has, hashNode := h.getNode(hashed)
	if !has {
		return false, nil
	}
	return true, hashNode.node.Data
}

func (h *HashMap) hashedKey(key string) int64 {
	return h.hashFn(key)
}

func (h *HashMap) getNode(key int64) (bool, *HashNode) {
	node := h.index[key]
	if nil == node {
		return false, nil
	}
	return true, node
}
