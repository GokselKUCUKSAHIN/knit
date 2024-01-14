package image

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func Test() {
	file, err := os.Open("/Users/gokselkucuksahin/GolandProjects/knit/img/norman.png")
	if err != nil {
		log.Fatal(err)
	}
	// IIFE
	defer func(catFile *os.File) {
		err := catFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	cat, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	size := cat.Bounds().Size()
	var w, h = size.X, size.Y
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			pixel := cat.At(x, y)
			fmt.Println(pixel)
		}
	}
}
