package reviews

func ValidParentheses(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
	
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
	
			if (top != '(' && char == ')') || 
			   (top != '[' && char == ']') || 
			   (top != '{' && char == '}') {
				return false
			}
		}
	}
	return len(stack) == 0
}