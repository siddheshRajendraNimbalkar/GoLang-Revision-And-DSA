package largestnumber

import "fmt"

func LargestNumber(nums ...int) {
	if len(nums) == 0 {
		fmt.Println("No numbers provided")
		return
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	fmt.Println(max)
}
