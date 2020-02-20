package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num_before =make([]int,0,0)
	var count int
	for i:=0;;i++{
		var num string
		_, _ = fmt.Scanf("%s",&num)
		if num=="" {break}
		v,_:=strconv.Atoi(num)
		num_before =append(num_before,v)
		count++
	}
	//fmt.Println(maopao(num_before,count))
	//var key int
	//_, _ = fmt.Scanf("%d",&key)
	//fmt.Println(erfen(num_before,count,key))
	qsort(num_before, 0, count-1)
	fmt.Println(num_before )
}

//冒泡排序，时间复杂度O(n^2),i循环了n次，j循环了n-j+1次
func maopao(num_after []int,count int)[]int{
	for i := 0; i < count - 1; i++ {
		for j := 0; j < count - 1 - i; j++ {
			if num_after[j] > num_after[j + 1]{
				num_after[j], num_after[j + 1] = num_after[j + 1], num_after[j]
			}
		}
	}
	return num_after
}

//二分查找，时间复杂度O(logn)脚标是2打不出来，设它的循环次数是t,它的循环条件2^t<n
func erfen(num_after []int,count int,key int)int {
	i:=0;j:=count-1;s:=0
	for {
		m:=(i+j)/2
		if num_after[m]==key{
			s=m;break
		} else if key<num_after[m]{
			j=m-1
		}else{ i=m+1}
		if i>j {
			break
		}
	}
	return s
}

//快速排序,时间复杂度O(nlogn)
func qsort(arr []int, first, last int) {
	flag := first
	left := first
	right := last
	if first >= last {
		return
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for first < last {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		for first < last {
			if arr[last] >= arr[flag] {
				last--
				continue
			}
			// 交换数据
			arr[last], arr[flag] = arr[flag], arr[last]
			flag = last
			break
		}
		for first < last {
			if arr[first] <= arr[flag] {
				first++
				continue
			}
			arr[first], arr[flag] = arr[flag], arr[first]
			flag = first
			break
		}
	}
	qsort(arr, left, flag-1)
	qsort(arr, flag+1, right)
}