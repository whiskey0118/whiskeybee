package main

import "fmt"

func main() {

	//b := make([]byte,10)
	//_,err := rand.Read(b)
	//if err != nil {
	//	fmt.Println("error",err)
	//	return
	//}
	//fmt.Println(bytes.Equal(b,make([]byte,10)))
	a := fmt.Sprintf("SELECT username,email FROM user where username='%s' or email='%s'", "whiskey", "funny")
	fmt.Println(a)

}
