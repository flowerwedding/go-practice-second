package main

import (
	"fmt"
	"strconv"
)

func main(){
	var slice_read =make([]int,0,0)
	var llen int
	for i:=0;;i++{
		var x string
		_, _ = fmt.Scanf("%s",&x)
		if x=="" {break}
		v, _ :=strconv.Atoi(x)
		slice_read=append(slice_read,v)
		llen++
	}
	slice_read=sort(slice_read,llen)
	//fmt.Println(slice_read)
	var slice_need =make([]int,llen)
	p:=0;q:=1
	for i:=0;i<=llen-1;i++{
		if slice_read[i]%2==0{slice_need[p]=slice_read[i];p+=2} else
		{slice_need[q]=slice_read[i];q+=2}
	}
	fmt.Println(slice_need)
}

func sort(num_after []int,count int)[]int{
	for i := 0; i < count - 1; i++ {
		for j := 0; j < count - 1 - i; j++ {
			if num_after[j] > num_after[j + 1]{
				num_after[j], num_after[j + 1] = num_after[j + 1], num_after[j]
			}
		}
	}
	//fmt.Println("num_aftr:",num_after)
	return num_after
}