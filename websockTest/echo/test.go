package main

import (
	"fmt"
	"time"
)

func main() {
	a, _ := time.Parse("2006-01-02", "2020-06-18")
	b := a.AddDate(0, 0, 1)
	fmt.Println(b)
	fmt.Println(a)
}

func test(s string, t string) bool {

	//s1 :=0
	//t1 :=0
	//for t1<len(t)&&s1<len(s){
	//	if s[s1]==t[t1]{
	//		s1++
	//	}
	//	t1++
	//}
	//
	//if s1==len(s){
	//	return true
	//}
	//return false

	var t1, s1 int
	for t1 < len(t) && s1 < len(s) {
		if t[t1] == s[s1] {
			s1++
		}
		t1++
	}
	fmt.Println(s1)
	fmt.Println(t1)
	return t1 == s1
}
