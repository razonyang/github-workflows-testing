package main

import (
	"fmt"
	"os"
)

func main() {
	dir, dirErr := os.UserCacheDir()
	fmt.Println(dir, dirErr)
}
