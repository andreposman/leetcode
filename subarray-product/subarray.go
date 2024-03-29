package subarray_product

import "fmt"

func Run() {
	/*
		nums = [10,5,2,6], k = 100
		   Steps:
		    - Get len(nums)
		    - Get nuns content by position
		    - Calculate the product of content
		    - Check if is less than k
		    - return the count of arr that the product is less than k
	*/

	var prevPosition int
	var subArr []int
	nums := []int{10, 5, 2, 6}
	totalProduct := 1
	k := 100

	fmt.Println("nums:", nums)
	fmt.Println("k:", k)

	for i := 0; i <= len(nums)-1; i++ {
		prevPosition = nums[i]
		fmt.Printf("\nCounter: %d | nums[i] = %d | Total Product: %d | Prev. Position: %d", i, nums[i], totalProduct, prevPosition)
		//fmt.Println("Counter: ", i)
		//fmt.Println("Nums[i]: ", nums[i])
		//fmt.Println("Previous Num[i]: ", prevPosition)
		if i == 0 {
			prevPosition = nums[0]
		}

		totalProduct = totalProduct * prevPosition

		if totalProduct < k {
			fmt.Printf("\nTotal Product: %d is lower than %d.\n", totalProduct, k)
			subArr = append(subArr, nums[i])
			//	TODO: find out how to push each pos into arr until < k
		} else {
			fmt.Printf("\nTotal Product: %d is HIGHER than %d.\n", totalProduct, k)
		}

		fmt.Println("Sub-array Count: ", subArr)
	}
}
