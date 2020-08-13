package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	fmt.Printf("%x  %x\n", md5.Sum([]byte("what")), "我")
	fmt.Printf("%x", "what")

}

//func thirdMax(nums []int) int {
//	if len(nums) >3 {
//		if nums[0] <nums[1] {
//			return nums[1]
//		}
//	}
//
//
//
//}
