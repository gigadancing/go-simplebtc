package hmap

// 实现hashmap原理

//
var buckets = make([]*Node, 16)

func InitBuckets() {
	for i := 0; i < 16; i++ {
		buckets[i] = CreateHead(KV{Key: "head", Value: "node"})
	}
}

// 向hashmap中保存键值对
func AddKV(k string, v string) {
	index := hashCode(k)            // 计算key的散列值
	head := buckets[index]          // 在数组中获得头结点
	tail := Tail(head)              // 获得尾结点
	Add(KV{Key: k, Value: v}, tail) // 添加节点
}

// 获取键值对
func GetValue(key string) string {
	index := hashCode(key)
	head := buckets[index]

	return FindValueByKey(key, head)
}

func hashCode(key string) int {
	var index = 0
	index = int(key[0])
	for k := 0; k < len(key); k++ {
		index *= 1103515245 + int(key[k])
	}
	index >>= 27
	index &= 16 - 1

	return index
}
