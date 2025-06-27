package main

import (
	"fmt"

	helloworld "github.com/siddheshRajendraNimbalkar/GoLang-Revision-And-DSA/Revision/Beginner/HelloWorld"
	swapvar "github.com/siddheshRajendraNimbalkar/GoLang-Revision-And-DSA/Revision/Beginner/SwapVar"
)

func Beginner() {
	helloworld.HelloWorld()
	var x interface{} = 100
	var y interface{} = 500

	swapvar.SwapVar(&x, &y)

	fmt.Println("x: ", x, "y: ", y)
}

func main() {
	Beginner()
}
