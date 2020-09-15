package main

import "fmt"

type evitionAlgo interface {
	evict(c *cache)
}
type cache struct {
	storage     map[string]string
	evitionAlgo evitionAlgo
	capacity    int
	maxCapacity int
}
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

func initCache(e evitionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:     storage,
		evitionAlgo: e,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *cache) setEvictionAlgo(e evitionAlgo) {
	c.evitionAlgo = e
}
func (c *cache) evict() {
	c.evitionAlgo.evict(c)
	c.capacity--
}
func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}
func (c *cache) get(key string) {
	delete(c.storage, key)
}
func main() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")
	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
