package factorialwithrecursion

import "fmt"

func factorialRecursion(num int64) int64 {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorialRecursion(num-1)
}

func FactorialRecursionAns(num int64) {
	ans := factorialRecursion(num)
	fmt.Printf("Factorial of %v is: %v", num, ans)
}
