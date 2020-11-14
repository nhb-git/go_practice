package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// 声明并初始化带缓冲的读取器
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input you name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("Fount an error : %s\n", err)
	} else {
		fmt.Printf("hello, %s", input)
	}
}