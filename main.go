package main

import (
	"fmt"

	structure "github.com/siddheshRajendraNimbalkar/GoLang-Revision-And-DSA/Revision/Intermediate/Structure"
)

func Beginner() {
	// Hello World
	// helloworld.HelloWorld()

	// Swap Two Variable
	// var x interface{} = 100
	// var y interface{} = 500

	// swapvar.SwapVar(&x, &y)

	// fmt.Println("x: ", x, "y: ", y)

	// Even Odd
	// evenodd.EvenOdd(2)
	// evenodd.EvenOdd(1)
	// evenodd.EvenOdd(0)

	// Largest Number
	// largestnumber.LargestNumber(1, 8, 3)

	// calculator
	// a := calculator.NewCalculatorImplementation(2, 3)
	// a.Addition()
	// a.Substraction()
	// a.Multiplication()
	// a.Dividion()

	// Factorial With Recursion
	// factorialwithrecursion.FactorialRecursionAns(6)

	// Prime Number
	// nums := 223
	// fmt.Println("Prime Number ", nums, ": ", primenumber.Prime(int64(nums)))
}

func Intermediate() {
	Car := structure.NewCarStruct("Toyota", "Camry", 2022)
	fmt.Println(Car.PrintCarDetails())
}

func main() {
	// Beginner()
	Intermediate()
}
