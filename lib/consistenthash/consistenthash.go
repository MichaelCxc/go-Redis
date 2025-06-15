package consistenthash

import (
	"hash/crc32"
	"sort"
)

type HashFunc func(data []byte) uint32

type NodeMap struct {
	hashFunc      HashFunc
	nodeHashValue []int
	nodeHashMap   map[int]string
}

func NewNodeMap(fn HashFunc) *NodeMap {
	m := &NodeMap{
		hashFunc:    fn,
		nodeHashMap: make(map[int]string),
	}

	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}

	return m
}

func (m *NodeMap) IsEmpty() bool {
	return len(m.nodeHashValue) == 0
}

func (m *NodeMap) AddNode(keys ...string) {
	for _, key := range keys {
		if key == "" {
			continue
		}
		hash := int(m.hashFunc([]byte(key)))
		m.nodeHashValue = append(m.nodeHashValue, hash)
		m.nodeHashMap[hash] = key
	}
	sort.Ints(m.nodeHashValue)
}

func (m *NodeMap) PickNode(keys string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := int(m.hashFunc([]byte(keys)))
	idx := sort.Search(len(m.nodeHashValue), func(i int) bool {
		return m.nodeHashValue[i] >= hash
	})
	if idx == len(m.nodeHashValue) {
		idx = 0
	}
	return m.nodeHashMap[m.nodeHashValue[idx]]
}
