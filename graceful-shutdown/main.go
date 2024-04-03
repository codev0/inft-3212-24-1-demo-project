package main

import (
	"fmt"
	"os"
)

func main() {
	// if err := simpleServe(); err != nil {
	// 	fmt.Println("err:", err)
	// 	os.Exit(1)
	// }

	if err := fullServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
