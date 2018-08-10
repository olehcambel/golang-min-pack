package main

import (
	"fmt"
)

type geek struct {
	ign, skill string
	id         int
}

func (y *geek) getData() {
	fmt.Println("important news:")
}

func main() {
	cambel := geek{ign: "olehcambel", id: 57322911}

	cambel.getData()

	fmt.Println(cambel)
}
