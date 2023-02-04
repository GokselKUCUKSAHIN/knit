package main

import (
	"fmt"
	"knit/knit"
)

func main() {
	// fmt.Println(knit.GeneratePinList(400, 600, 600))
	arr := knit.GetPointListOnLine(300, 0, 465, 49)
	fmt.Println(arr)
	fmt.Println(len(arr))
}
