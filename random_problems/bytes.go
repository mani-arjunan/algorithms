package main

import "fmt"

func practiceBytes() {
	str := "È"
	str2 := "Ã"

	by := []byte(str)

	fmt.Println(string(by[0]))

	fmt.Println(string([]byte(str2)))
	fmt.Println(string([]byte(str2)[0]))

	// for _, val := range []byte(str2) {
	// 	fmt.Println(val)
	// }
	// for _, val := range str2 {
	// 	fmt.Println(val)
	// }
}
