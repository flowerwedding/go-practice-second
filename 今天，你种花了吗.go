package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

花园共有v个位置，已种树u棵，将种树n棵
当v为偶数时，n<=v\2-u
当v为奇数时，如果树都在奇数位上，n<=(v+1)\2-u
             否则，n<(v+1)\2-u
奇数乘奇数等于奇数

*/

func main(){
    reader := bufio.NewReader(os.Stdin)
    data, _, _ := reader.ReadLine()
    mult := 1
     v ,u := 0 ,0
    position:=strings.Index(string(data),"[")
	positionlast:=strings.Index(string(data),"]")
    for i:=position;i<positionlast;i+=2{
		t, _ := strconv.Atoi(string(data[i+1]))
    	v++
    	if t ==1 {
    		u++
    		mult*=(i-position+2)/2
    	}
	}
	position_n :=strings.Index(string(data),"n")
	n, _ := strconv.Atoi(string(data[position_n+4:]))
	if v%2 == 0 {
		if n <= v/2-u {
			fmt.Println("true")
		}else {
			fmt.Println("false")
		}
	}else {
		if mult%2 ==1 {
			if n <= (v+1)/2-u {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		}else {
			if n < (v+1)/2-u{
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		}
	}
}
