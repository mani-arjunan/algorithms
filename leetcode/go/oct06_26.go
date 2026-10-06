package main

func minAdditionToMakeValidString(str string) int {
	var stack []string
	var closeStack []string

	for _, val := range str {
		if string(val) == "(" {
			stack = append(stack, "(")
		} else {
			if len(stack) == 0 {
				closeStack = append(closeStack, ")")
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(closeStack) + len(stack)
}

func main() {
	str := "(("
	minAdditionToMakeValidString(str)
}
