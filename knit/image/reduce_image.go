package image

func ReduceImage() []uint8 {

	// TODO Take rgb image and return graystace uint8 array
	return nil
}

func RgbToGray(r, g, b uint8) uint8 {
	// GrayScale = 0.299R + 0.587G + 0.114B
	result := uint8(float64(r)*0.299) + uint8(float64(g)*0.587) + uint8(float64(b)*0.114)
	if result > 255 {
		return 255
	}
	return result
}
