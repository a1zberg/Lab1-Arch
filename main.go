package main

import (
	//"bufio"
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
	fmt.Println("Version is: " + BuildVersion)
	//if input, err := os.Open("version.txt"); err == nil {
	//	defer input.Close()
	//	scanner := bufio.NewScanner(input)
	//	for scanner.Scan() {
	//		buildVersion := scanner.Text()
	//		fmt.Println("Version is: " + buildVersion)
	//	}
	//}
}
