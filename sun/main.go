package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("---------------------------------------------------------")

	var arr = [4]int{1, 2, 3, 4}
	const me int = 657
	var words = []string{}
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for num := range nums {
		words = append(words, "sad")
		fmt.Println(words, num)
	}
	fmt.Println("---------------------------------------------------------")

	fmt.Println(words)

	fmt.Println("---------------------------------------------------------")
	// for i := 0; i < me; i++ {
	// 	fmt.Println("This number is ", i)
	// }

	if me != 200 {
		return
	}
	fmt.Println("---------------------------------------------------------")

	fmt.Println(nums, me)
	fmt.Println(arr)
}
