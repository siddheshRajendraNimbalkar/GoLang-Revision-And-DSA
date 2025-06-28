// Swap Two Variables

// Swap two variables without using a third variable.

package swapvar

import (
	"fmt"
	"log"
	"reflect"
)

func SwapVar(a, b *interface{}) {
	if reflect.TypeOf(*a) != reflect.TypeOf(*b) {
		log.Panicln("enter same data type")
	} else {
		fmt.Println("a: ", *a, "b: ", *b)
		*a, *b = *b, *a
		fmt.Println("a: ", *a, "b: ", *b)
	}

}
