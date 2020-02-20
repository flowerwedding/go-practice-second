package main

import ("fmt"
         "math")

func main() {
	var n int
	_, _ = fmt.Scanf("%d",&n)
	//fibN := fib1(n)
	fibN := fib2(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib1(x int) int {//时间复杂度O(n)
	if x < 2 {
		return x
	}
	return fib1(x-1) + fib1(x-2)
}

func fib2(x int) int{//时间复杂度O(1)
	ans:=(1/math.Sqrt(5))*(math.Pow((1+math.Sqrt(5))/2, float64(x))+math.Pow((1-math.Sqrt(5))/2, float64(x)))
	return int(ans)
}