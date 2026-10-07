package main

import "fmt"

func practiceBytes() {
	str := "È"
	str2 := "Ã"

	by := []byte(str)

	fmt.Println(string(by[0]))

	fmt.Println(string([]byte(str2)))
	fmt.Println(string([]byte(str2)[0]))

	// converts a number into its associated unicode character
	fmt.Println(string(123))

	// converts array of bytes into its associated character based on utf8 encoder, multiple bytes
	// for one single character are common in emojis, other languages like latin, tamil etc
	fmt.Println(string([]byte{240, 159, 152, 128}))

	// for _, val := range []byte(str2) {
	// 	fmt.Println(val)
	// }
	// for _, val := range str2 {
	// 	fmt.Println(val)
	// }
}
