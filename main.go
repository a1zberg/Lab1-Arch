package main

import (
	"fmt"
	"os"

	"./implementation"
)

func main() {
	var expression string
	fmt.Println("Введите выражение в префиксной форме:")
	fmt.Fscan(os.Stdin, &expression)
	result := implementation.PrefToPost(expression)
	fmt.Println("Это же число в постфиксной форме:")
	fmt.Println(result)
}
