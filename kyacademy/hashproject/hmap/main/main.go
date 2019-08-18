package main

/* Hash
 * 1.可以将不同长度明文转换成相同长度的密文
 *   abcd ->  avasfsa
 *   ab   ->  asdfeer
 */

/* 一个优秀的hash算法：
 * 1.正想快速
 *   给定明文和hash算法，在有限时间和有限资源内能计算出hash值
 *
 * 2.逆向困难
 *   给定（若干）hash值，在有限时间内很难（基本不可能）逆推出明文
 *
 * 3.输入敏感
 *   原始输入信息修改一点信息，产生的hash值看起来应该有很大不同
 *
 * 4.冲突避免
 *   很难找到两端内容不同的明文，是的它们的hash值一致（发生冲突）。即对于任意两个不同的数据块，其hash值相同的可能性极小；
 *   对于一个给定的数据块，找到它们hash值相同的数据值极为困难。
 */

// 简单的hash散列函数
func test(a int) int {
	return (a+2+(a<<1))%8 ^ 5
}

// 将任何长度的字符串通过运算，散列成0-15的整数
// 通过hashCode散列出0-15的数字的概率是相等的
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

func main() {

}
