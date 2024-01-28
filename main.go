package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hi")
	var port = os.Getenv("PORT")
	fmt.Println(port)
}
