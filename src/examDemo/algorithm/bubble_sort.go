package main

import "fmt"

// 冒泡排序
// 时间复杂度： O(n的平方)
func main() {

	data := []int{12,31,14,51,16,33,0,9,17,25}

	bubbleSort(data)
	// fmt.Println(data)   // [0 9 12 14 16 17 25 31 33 51]

	// 算法优化：
	/*  不加标识位： (排序过程如下, 后面几趟没有进行任何数据交换，但是依然进行了好几趟对比，浪费资源)

			[12 14 31 16 33 0 9 17 25 51]
			[12 14 16 31 0 9 17 25 33 51]
			[12 14 16 0 9 17 25 31 33 51]
			[12 14 0 9 16 17 25 31 33 51]
			[12 0 9 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]


		加了标识位 (排序过程如下： 减少了对比次数)：

	        [12 14 31 16 33 0 9 17 25 51]
			[12 14 16 31 0 9 17 25 33 51]
			[12 14 16 0 9 17 25 31 33 51]
			[12 14 0 9 16 17 25 31 33 51]
			[12 0 9 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]
			[0 9 12 14 16 17 25 31 33 51]


	*/
}

func bubbleSort(data []int){

	// 循环 len(data)-1 趟
	for i:=0;i<len(data)-1;i++ {
		// 加入一个标志位，标识 是否有数据交换
		exchange := false

		// 无序区比较 ([0, len(data)-i-1)
		for j:=0;j<len(data)-i-1;j++ {
			if data[j] > data[j+1] {
				exchange = true
				data[j],data[j+1] = data[j+1],data[j]
			}
		}
		fmt.Println(data)
		// 如果 一趟比较 没有进行任何数字交换，则代表已经是有序的
		if exchange == false {
			return
		}
	}
}