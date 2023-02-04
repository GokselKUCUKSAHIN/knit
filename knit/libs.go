package knit

import "math"

const (
	Half    float64 = 0.5
	Quarter float64 = 0.25
	Pi              = math.Pi
	TwoPi           = Pi * 2
	HalfPi          = Pi * Half
)

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func floatEquals(first, second float64) bool {
	return math.Abs(first-second) < 1e-7
}

func GeneratePinList(length, w, h int) []int16 {
	width := float64(w)
	height := float64(h)
	centerX := width * Half
	centerY := height * Half
	radius := centerX
	angleUnit := TwoPi / float64(length)
	pins := make([]int16, length<<1, length<<1)
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
		pins[i<<1] = int16(x)
		pins[i<<1+1] = int16(y)
	}
	return pins
}

func IsDotOnTheLine(dotX, dotY, startX, startY, endX, endY float64) bool {
	if floatEquals(endX, startX) {
		return floatEquals(dotX, endX)
	}
	slope := (endY - startY) / (endX - startX)
	intercept := startY - slope*startX
	blockTopY := dotY + Half
	blockBottomY := dotY - Half
	blockLeftY := slope*(dotX-Half) + intercept
	blockRightY := slope*(dotX+Half) + intercept
	if math.Abs(slope) <= 1 {
		return (blockLeftY >= blockBottomY && blockLeftY <= blockTopY) ||
			(blockRightY >= blockBottomY && blockRightY <= blockTopY)
	} else if slope > 0 {
		return !(blockLeftY > blockTopY || blockRightY < blockBottomY)
	} else {
		return !(blockRightY > blockTopY || blockLeftY < blockBottomY)
	}
}

func GetPointListOnLine(startX, startY, endX, endY int16) []int16 {
	pointList := make([]int16, 0, 843)
	var movementX, movementY float64
	if endX > startX {
		movementX = 1
	} else {
		movementX = -1
	}
	if endY > startY {
		movementY = 1
	} else {
		movementY = -1
	}
	var currentX, currentY, startXf, startYf, endXf, endYf = float64(startX), float64(startY), float64(startX), float64(startY), float64(endX), float64(endY)
	for i := int16(0); (currentX != endXf || currentY != endYf) && i < 1000; i++ {
		pointList = append(pointList, int16(currentX))
		pointList = append(pointList, int16(currentY))
		if IsDotOnTheLine(currentX+movementX, currentY, startXf, startYf, endXf, endYf) {
			currentX += movementX
		} else {
			currentY += movementY
		}
	}
	pointList = append(pointList, int16(endX))
	pointList = append(pointList, int16(endY))
	return pointList
}
