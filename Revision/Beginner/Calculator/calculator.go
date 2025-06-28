// Implement a simple calculator that performs addition, subtraction, multiplication, and division.

package calculator

import "fmt"

type Calculator interface {
	Addition()
	Substraction()
	Multiplication()
	Dividion()
}

type CalculatorImplementation struct {
	a float64
	b float64
}

func NewCalculatorImplementation(a float64, b float64) *CalculatorImplementation {
	return &CalculatorImplementation{
		a: a,
		b: b,
	}
}

func (c CalculatorImplementation) Addition() {
	fmt.Println("Addition is: ", c.a+c.b)
}

func (c CalculatorImplementation) Substraction() {
	fmt.Println("Substraction is: ", c.a-c.b)
}

func (c CalculatorImplementation) Multiplication() {
	fmt.Println("Multiplication is: ", c.a*c.b)
}

func (c CalculatorImplementation) Dividion() {
	if c.a == 0 || c.b == 0 {
		panic("cannot divide by zero")
	}
	fmt.Println("Dividion is: ", c.a/c.b)

}
