package slice

import "fmt"

func SliceCustom() {

	num_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(num_slice)

	numbers := make([]int, 3, 5)
	numbers = append(numbers, 5, 6, 3, 445, 6)
	fmt.Println(numbers)

	out_len := len(numbers)
	fmt.Println("Slice length", out_len)
	out_cap := cap(numbers)
	fmt.Println("Slice capacity", out_cap)
}
