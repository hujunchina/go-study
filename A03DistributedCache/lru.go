package A03DistributedCache
//https://geektutu.com/post/geecache-day1.html
//LRU（least recently used）最近最少使用结合了时间和频率的缓存删除算法

import(
	"container/list"
)

type Cache struct {
	maxBytes int64
	nBytes int64
	//双向链表list.List
	ll *list.List
	//键是字符串，值是双向链表中对应节点的指针
	cache map[string]*list.Element
	//清除键值时触发，某条记录被移除时的回调函数，可以为 nil
	OnEvicted func(key string, value Value)
}

type entry struct{
	key string
	value Value
}
//值是实现了 Value 接口的任意类型
type Value interface {
	//返回值所占用的内存大小
	Len() int
}

// Cache类的构造方法
// 传入最大节点值，回调方法
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

//从字典中找到对应的双向链表的节点
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok{
		//将该节点移动到队尾,约定 front 为队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

//这里的删除，实际上是缓存淘汰。即移除最近最少访问的节点（队首）
func (c *Cache) RemoveOldest() {
	//取到队首节点
	ele := c.ll.Back()
	if ele != nil{
		//从链表中删除
		c.ll.Remove(ele)
		//从map中删除
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		//更新内存占用
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		//调用回调函数
		if c.OnEvicted != nil{
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

//新增和修改
func (c *Cache) Add(key string, value Value){
	if ele, ok := c.cache[key]; ok{
		//移到队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		//插入队尾
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	//超过了设定的最大值 c.maxBytes移除最少访问的节点
	for c.maxBytes!=0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
