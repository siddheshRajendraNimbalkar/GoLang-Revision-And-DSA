package main

import calculator "github.com/siddheshRajendraNimbalkar/GoLang-Revision-And-DSA/Revision/Beginner/Calculator"

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
	a := calculator.NewCalculatorImplementation(2, 3)
	a.Addition()
	a.Substraction()
	a.Multiplication()
	a.Dividion()
}

func main() {
	Beginner()
}
