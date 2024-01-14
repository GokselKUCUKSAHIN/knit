package high_level

import (
	"knit/knit/high_level/model"
	"knit/knit/utils"
	"math"
)

const (
	Half    float64 = 0.5
	Quarter float64 = 0.25
	Pi              = math.Pi
	TwoPi           = Pi * 2
	HalfPi          = Pi * Half
)

func GeneratorPinList(length int, width, height float64) []*model.Vector {
	centerX := width * Half
	centerY := height * Half
	radius := centerX
	angleUnit := TwoPi / float64(length)
	pins := make([]*model.Vector, 0, 2*length)
	for i := 0; i < length; i++ {
		angle := angleUnit*float64(i) - HalfPi
		x := math.Round(centerX + radius*math.Cos(angle))
		y := math.Round(centerY + radius*math.Sin(angle))
		if x >= width {
			x = width - 1
		}
		if y >= height {
			y = height - 1
		}
		pins = append(pins, model.CreateVector(x, y))
	}
	return pins
}

func IsDotOnTheLine(dot, start, end *model.Vector) bool {
	if utils.FloatEquals(start.X, end.X) {
		return utils.FloatEquals(dot.X, end.X)
	}
	slope := start.Slope(end)
	intercept := float64(start.Y) - slope*float64(start.X)
	blockTopY := float64(dot.Y) + Half
	blockBottomY := float64(dot.Y) - Half
	blockLeftY := slope*(float64(dot.X)-Half) + intercept
	blockRightY := slope*(float64(dot.X)+Half) + intercept
	if math.Abs(slope) <= 1 {
		return (blockLeftY >= blockBottomY && blockLeftY <= blockTopY) ||
			(blockRightY >= blockBottomY && blockRightY <= blockTopY)
	} else if slope > 0 {
		return !(blockLeftY > blockTopY || blockRightY < blockBottomY)
	} else {
		return !(blockRightY > blockTopY || blockLeftY < blockBottomY)
	}
}

func GetPointListOnTheLine(start, end *model.Vector) []*model.Vector {
	pointList := make([]*model.Vector, 0)
	var movementX, movementY int16
	if end.X > start.X {
		movementX = 1
	} else {
		movementX = -1
	}
	if end.Y > start.Y {
		movementY = 1
	} else {
		movementY = -1
	}
	currentX := start.X
	currentY := start.Y
	for i := int16(0); (currentX != end.X || currentY != end.Y) && i < 1000; i++ {
		pointList = append(pointList, model.CreateVector(currentX, currentY))

		if IsDotOnTheLine(model.CreateVector(currentX+movementX, currentY), start, end) {
			currentX += movementX
		} else {
			currentY += movementY
		}
	}
	pointList = append(pointList, end)
	return pointList
}

func ReduceImageData(image *model.Image, start, end *model.Vector) {
	dotList := GetPointListOnTheLine(start, end)
	for _, dot := range dotList {
		image.Inc(dot.X, dot.Y, 50)
	}
}

func GetLineScore(image *model.Image, start, end *model.Vector) float64 {
	dotList := GetPointListOnTheLine(start, end)
	sum := 0
	for _, dot := range dotList {
		sum += int(255 - image.At(dot.X, dot.Y))
	}
	return float64(sum) / float64(len(dotList)*255)
}

func IsLineDrawn() {

}
