package primenumber

import "fmt"

func Prime(nums int64) bool {
	if nums <= 1 {
		fmt.Println("It is The Prime Number")
		return false
	}
	for i := int64(2); i*i <= nums; i++ {
		if nums%i == 0 {
			return false
		}
	}
	return true
}
