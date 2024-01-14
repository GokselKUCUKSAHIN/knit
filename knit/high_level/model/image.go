package model

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

type Image struct {
	Path  string
	Slice []uint8
	w     int16
	h     int16
	data  image.Image
}

func NewImage(path string) (*Image, error) {
	img, err := LoadImage(path)
	if err != nil {
		return nil, err
	}
	slice, w, h := ImageToUInt8(img)
	return &Image{
		Path:  path,
		Slice: slice,
		w:     w,
		h:     h,
		data:  img,
	}, nil
}

func (base *Image) GetData() image.Image {
	return base.data
}

func (base *Image) Bounds() (int16, int16) {
	return base.w, base.h
}

func (base *Image) GetIndex(x, y int16) int16 {
	w, _ := base.Bounds()
	return w*y + x
}

func (base *Image) At(x, y int16) uint8 {
	return base.Slice[base.GetIndex(x, y)]
}

func (base *Image) Set(x, y int16, newValue uint8) uint8 {
	result := CapUint8(int16(newValue))
	base.Slice[base.GetIndex(x, y)] = result
	return result
}

func (base *Image) Inc(x, y int16, amount int16) uint8 {
	prev := int16(base.At(x, y))
	return base.Set(x, y, CapUint8(prev+amount))
}

func (base *Image) Dec(x, y int16, amount int16) uint8 {
	return base.Inc(x, y, -amount)
}

func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func RgbaToGray(color color.Color) uint8 {
	// 0.299R + 0.587G + 0.114B
	r, g, b, _ := color.RGBA()
	result := uint8(float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114)
	if result > 255 {
		return 255
	}
	return result
}

func ImageToUInt8(src image.Image) ([]uint8, int16, int16) {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := make([]uint8, w*h, w*h)
	for y := 0; y < h; y++ {
		row := y * w
		for x := 0; x < w; x++ {
			result[row+x] = RgbaToGray(src.At(x, y))
		}
	}
	return result, int16(w), int16(h)
}

func CapUint8(value int16) uint8 {
	result := value
	if value > 255 {
		result = 255
	} else if result < 0 {
		value = 0
	}
	return uint8(result)
}
