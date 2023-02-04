package high_level

import (
	"knit/knit/high_level/model"
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
		if x > width {
			x -= 1
		}
		if y > height {
			y -= 1
		}
		pins = append(pins, model.CreateVector(x, y))
	}
	return pins
}

// dotX, dotY, startX, startY, endX, endY float64

//func IsDotOnTheLine(dot, start, end *model.Vector) bool {
//	if utils.FloatEquals(start.X, end.X) {
//		return utils.FloatEquals(dot.X, end.X)
//	}
//	slope := start.Slope(end)
//	intercept := float64(start.Y) - slope*float64(start.X)
//	blockTopY := dot.Y + Half
//	blockBottomY := dotY - Half
//	blockLeftY := slope*(dotX-Half) + intercept
//	blockRightY := slope*(dotX+Half) + intercept
//	if math.Abs(slope) <= 1 {
//		return (blockLeftY >= blockBottomY && blockLeftY <= blockTopY) ||
//			(blockRightY >= blockBottomY && blockRightY <= blockTopY)
//	} else if slope > 0 {
//		return !(blockLeftY > blockTopY || blockRightY < blockBottomY)
//	} else {
//		return !(blockRightY > blockTopY || blockLeftY < blockBottomY)
//	}
//}
