package main

func removeOuterParanthesis(str string) string {
	stack := make([]string, 0)
	var result string

	for _, val := range str {
		if val == '(' {
			stack = append(stack, string(val))

			if len(stack) > 1 {
				result += "("
			}
		} else {
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				result += ")"
			}
		}
	}

	return result

}
