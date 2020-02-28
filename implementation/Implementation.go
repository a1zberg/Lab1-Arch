package implementation

import (
	"fmt"
	"strings"
)
import (
	"github.com/golang-collections/collections/stack"
)

func isOperator(x string) bool {
	switch x {
	case "+":
		return true
	case "-":
		return true
	case "/":
		return true
	case "*":
		return true
	}
	return false
}

func PrefToPost(input string) string {

	stack := stack.New()
	length := len(input)
	if length == 0 {
		return "The string is empty"
	}
	if strings.Contains(input, "!") {
		return "Unacceptable symbols!"
	} //Для теста
	for i := (length - 1); i >= 0; i-- {
		if isOperator(string(input[i])) {
			// pop two operands from stack
			op1help := stack.Peek()
			stack.Pop()
			op1 := fmt.Sprintf("%v", op1help)
			op2help := stack.Peek()
			stack.Pop()
			op2 := fmt.Sprintf("%v", op2help)
			// concat the operands and operator
			temp := op1 + op2 + string(input[i])
			// Push String temp back to stack
			stack.Push(temp)
		} else {
			// push the operand to the stack
			stack.Push(string(input[i]))
		}
	}
	result := fmt.Sprintf("%v", stack.Peek())
	return result
}
