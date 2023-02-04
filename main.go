package main

import (
	"fmt"
	"knit/knit/low_level"
)

func main() {
	pinList := low_level.GeneratePinList(400, 600, 600)
	counts := make([]int, 0)
	for i := 0; i < len(pinList)/2; i++ {
		xA := pinList[2*i]
		yA := pinList[2*i+1]
		for j := 0; j < i; j++ {
			xB := pinList[2*j]
			yB := pinList[2*j+1]
			count := low_level.GetPointListOnLine(xA, yA, xB, yB)
			counts = append(counts, len(count))
		}
	}
	sum := 0
	for _, count := range counts {
		sum += count
	}
	fmt.Printf("sum: %d, count: %d, avg: %.2f\n", sum, len(counts), float64(sum)/float64(len(counts)))
}
