package main

import (
	"fmt"
)

type geek struct {
	ign, skill string
	id         int
}

func main() {
	cambel := geek{ign: "olehcambel", id: 57322911}
	// cambel := geek{"olehcambel", "React", 57322911}

	// var cambel geek
	// cambel.ign = "olehcambel"
	fmt.Println(cambel)
}
