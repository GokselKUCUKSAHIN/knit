package main

import (
	"fmt"
	"knit/knit/high_level/model"
	"log"
)

func main() {
	//pinList := high_level.GeneratorPinList(400, 600, 600)
	//pinS := pinList[201]
	//pinE := pinList[305]
	//fmt.Println(pinS)
	//fmt.Println(pinE)
	//pointList := high_level.GetPointListOnTheLine(pinS, pinE)
	//fmt.Println(pointList)
	//fmt.Println(len(pointList))
	// image.Test()
	// model.LoadImage("/Users/gokselkucuksahin/GolandProjects/knit/img/norman.png")
	norman := "/Users/gokselkucuksahin/GolandProjects/knit/img/norman.png"
	img, err := model.NewImage(norman)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(img.Slice[:200])
	fmt.Println(img.At(599, 599))

	//counts := make([]int, 0)
	//for i := 0; i < len(pinList)/2; i++ {
	//	xA := pinList[2*i]
	//	yA := pinList[2*i+1]
	//	for j := 0; j < i; j++ {
	//		xB := pinList[2*j]
	//		yB := pinList[2*j+1]
	//		count := low_level.GetPointListOnLine(xA, yA, xB, yB)
	//		counts = append(counts, len(count))
	//	}
	//}
	//sum := 0
	//for _, count := range counts {
	//	sum += count
	//}
	//fmt.Printf("sum: %d, count: %d, avg: %.2f\n", sum, len(counts), float64(sum)/float64(len(counts)))
}
