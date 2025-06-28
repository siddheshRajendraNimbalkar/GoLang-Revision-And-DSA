// Check Even or Odd

package evenodd

import "fmt"

func EvenOdd(a int) {
	if a == 0 {
		fmt.Println("The Num is Zero")
	} else if a%2 == 0 {
		fmt.Println("The Num is Even")
	} else {
		fmt.Println("The Num is Odd")
	}
}
