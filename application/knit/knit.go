package knit

import (
	"fmt"
	"knit2024/application/model_domain"
	"knit2024/application/model_response"
	"math"
)

func Create(image [][]uint8, pinCount, lineLimit int, width, height int16) *model_response.KnitResponse {
	pins := generatePins(pinCount, width, height)
	var lines = make(model_domain.Lines, 0)
	lines = draw(lines, image, pins, 0, lineLimit)
	return model_response.NewKnitResponse(pins, lines, width, height)
}

func generatePins(pinCount int, width, height int16) []*model_domain.Vector2[int16] {
	pins := make([]*model_domain.Vector2[int16], 0, pinCount)
	centerX, centerY := width/2, height/2
	radius := float64(width-5) * 0.5
	angleInterval := 2 * math.Pi / float64(pinCount)
	for i := 0; i < pinCount; i++ {
		angle := float64(i)*angleInterval - math.Pi*0.5
		x := int16(math.Round(radius*math.Cos(angle) + float64(centerX)))
		if x == width {
			x -= 1
		}
		y := int16(math.Round(radius*math.Sin(angle) + float64(centerY)))
		if y == height {
			y -= 1
		}
		pins = append(pins, &model_domain.Vector2[int16]{X: x, Y: y})
	}
	return pins
}

func draw(lines model_domain.Lines, image [][]uint8, pins []*model_domain.Vector2[int16], pinStartIndex int16, lineLimit int) model_domain.Lines {
	highestScore, pinEndIndex := 0.0, 0
	for i, pin := range pins {
		i16 := int16(i)
		if isLineDrawn(lines, pinStartIndex, i16) ||
			isPinTooClose(pins, pinStartIndex, i16) ||
			pinStartIndex == i16 {
			continue
		}
		score := calculateLineScore(image, pins[pinStartIndex], pin)
		if score > highestScore {
			highestScore = score
			pinEndIndex = i
		}
	}
	if len(lines) <= lineLimit {
		lines = append(lines, model_domain.NewPair(pinStartIndex, int16(pinEndIndex)))
		reduceImage(image, pins[pinStartIndex], pins[pinEndIndex])
		lines = draw(lines, image, pins, int16(pinEndIndex), lineLimit)
	}
	return lines
}

func isOnTheLine(dot, lineStart, lineEnd *model_domain.Vector2[int16]) bool {
	if lineStart.X == lineEnd.X {
		return dot.X == lineEnd.X
	}
	if lineStart.Y == lineEnd.Y {
		return dot.Y == lineEnd.X
	}
	slope := float64(lineEnd.Y-lineStart.Y) / float64(lineEnd.X-lineStart.X)
	intercept := float64(lineStart.Y) - slope*float64(lineStart.X)
	blockTopY, blockBottomY := float64(dot.Y)+0.5, float64(dot.Y)-0.5
	blockLeftY, blockRightY := slope*(float64(dot.X)-0.5)+intercept, slope*(float64(dot.X)+0.5)+intercept
	if math.Abs(slope) <= 1.0 {
		return (blockLeftY >= blockBottomY && blockLeftY <= blockTopY) ||
			(blockRightY >= blockBottomY && blockRightY <= blockTopY)
	} else if slope > 0 {
		return !(blockLeftY > blockTopY || blockRightY < blockBottomY)
	} else {
		return !(blockRightY > blockBottomY || blockLeftY < blockBottomY)
	}
}

var pointOnTheLineCache = make(map[string][]*model_domain.Vector2[int16])

func getPointsOnTheLine(lineStart, lineEnd *model_domain.Vector2[int16], limit int16) []*model_domain.Vector2[int16] {
	cacheKey := fmt.Sprintf("%d_%d;%d_%d", lineStart.X, lineStart.Y, lineEnd.X, lineEnd.Y)
	if cachedResult, exists := pointOnTheLineCache[cacheKey]; exists {
		return cachedResult
	}
	points := make([]*model_domain.Vector2[int16], 0)
	movementX := int16(-1)
	if lineEnd.X > lineStart.X {
		movementX = 1
	}
	movementY := int16(-1)
	if lineEnd.Y > lineStart.Y {
		movementY = 1
	}
	currentX, currentY := lineStart.X, lineStart.Y
	for i := 0; (currentX != lineEnd.X || currentY != lineEnd.Y) && i < 1_000; i++ {
		points = append(points, &model_domain.Vector2[int16]{
			X: currentX,
			Y: currentY,
		})
		if isOnTheLine(&model_domain.Vector2[int16]{X: currentX + movementX, Y: currentY}, lineStart, lineEnd) {
			next := currentX + movementX
			if next < 0 {
				next = 0
			} else if next >= limit {
				next = limit - 1
			}
			currentX = next
		} else {
			next := currentY + movementY
			if next < 0 {
				next = 0
			} else if next >= limit {
				next = limit - 1
			}
			currentY = next
		}
	}
	points = append(points, lineEnd)
	pointOnTheLineCache[cacheKey] = points
	return points
}

func reduceImage(image [][]uint8, lineStart, lineEnd *model_domain.Vector2[int16]) {
	points := getPointsOnTheLine(lineStart, lineEnd, int16(len(image)))
	for i := 0; i < len(points); i++ {
		reducePoint(image, points[i])
	}
}

func reducePoint(image [][]uint8, dot *model_domain.Vector2[int16]) {
	value := uint16(image[dot.Y][dot.X]) + 50
	if value < 256 {
		image[dot.Y][dot.X] = uint8(value)
	} else {
		image[dot.Y][dot.X] = 255
	}
}

func calculateLineScore(image [][]uint8, lineStart, lineEnd *model_domain.Vector2[int16]) float64 {
	points := getPointsOnTheLine(lineStart, lineEnd, int16(len(image)))
	scores := make([]float64, 0, len(points))
	for _, point := range points {
		color := image[point.Y][point.X]
		colorScore := 1.0 - float64(color)/255.0
		scores = append(scores, colorScore)
	}
	totalScore := 0.0
	for _, score := range scores {
		totalScore += score
	}
	result := totalScore / float64(len(scores))
	return result
}

func isPinTooClose(pins []*model_domain.Vector2[int16], pinStart, pinEnd int16) bool {
	pinLen := float64(len(pins))
	pinDistance := math.Abs(float64(pinEnd - pinStart))
	if pinDistance > pinLen*0.5 {
		pinDistance = pinLen - pinDistance
	}
	return pinDistance < 25
}

func isLineDrawn(lines model_domain.Lines, pinStart, pinEnd int16) bool {
	for _, line := range lines {
		lineStartPinIndex, lineEndPinIndex := line.First(), line.Second()
		if (lineStartPinIndex == pinStart && lineEndPinIndex == pinEnd) ||
			(lineStartPinIndex == pinEnd && lineEndPinIndex == pinStart) {
			return true
		}
	}
	return false
}
