package main

import (
	"context"
	"fmt"
)

func main() {
	contextParent := context.Background()
	ctx1 := context.WithValue(contextParent, "key1", "Hello World!")
	ctx2 := context.WithValue(ctx1, "key2", "Hello Guys!")
	ctx3 := context.WithValue(ctx2, "key3", "Hello Bois!")
	ctx4 := context.WithValue(contextParent, "key4", "Hello kids!")
	ctx5 := context.WithValue(ctx1, "key5", "Hello Then")
	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))
	fmt.Println(ctx3.Value("key1"))
	fmt.Println(ctx4.Value("key1"))
	fmt.Println(ctx4.Value("key4"))

	var input string
	fmt.Scanln(&input)
}
