package dylan_algo

//https://leetcode.cn/problems/lru-cache/

type LRUCache struct {
	capacity   int
	cacheMap   map[int]LinkNode
	head, tail *LinkNode
}
type LinkNode struct {
	key, value int
	prev, next *LinkNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]LinkNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.updateNewNode(&node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	if ok {
		node.value = value
	} else {
		node = LinkNode{
			key:   key,
			value: value,
		}
		if len(this.cacheMap) >= this.capacity {
			this.removeOld()
		}
		this.cacheMap[key] = node
		this.updateNewNode(&node)

	}

}

func (this *LRUCache) removeOld() {
	tailNode := this.tail
	prevTail := tailNode.prev

	delete(this.cacheMap, tailNode.key)
	if prevTail != nil {
		prevTail.next = nil
		this.tail = prevTail
		tailNode.prev = nil
	} else {
		this.tail = nil
		this.head = nil
	}
}

func (this *LRUCache) updateNewNode(node *LinkNode) {
	nodePre := node.prev
	nodeNext := node.next
	nodePre.next = nodeNext
	nodeNext.prev = nodePre
	tailNode := this.tail
	node.prev = nil
	node.next = tailNode
	tailNode.prev = node
	this.tail = tailNode
}
