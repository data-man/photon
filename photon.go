package photon

import (
	"image"
	"image/png"
	"math"
	"os"
)

type Image struct {
	ImageData *image.RGBA
	Width     int
	Height    int
}

func (image *Image) FillRect(startPoint, endPoint Point, color Color) *Image {
	// Todo: Check if it can really be drawn in the space of the image
	for x := min(startPoint.X, endPoint.X); x <= max(startPoint.X, endPoint.X); x++ {
		for y := min(startPoint.Y, endPoint.Y); y <= max(startPoint.X, endPoint.Y); y++ {
			image.ImageData.Set(x, y, color)
		}
	}
	return image
}

func (image *Image) StrokeRect(startPoint, endPoint Point, lineWidth int, color Color) *Image {
	for x := min(startPoint.X, endPoint.X); x <= max(startPoint.X, endPoint.X); x++ {
		for y := min(startPoint.Y, endPoint.Y); y <= max(startPoint.Y, endPoint.Y); y++ {
			if x <= min(startPoint.X, endPoint.X)+lineWidth || x >= max(startPoint.X, endPoint.X)-lineWidth {
				image.ImageData.Set(x, y, color)
			}
			if y <= min(startPoint.Y, endPoint.Y)+lineWidth || y >= max(startPoint.Y, endPoint.Y)-lineWidth {
				image.ImageData.Set(x, y, color)
			}
		}
	}
	return image
}

func (image *Image) StrokeLine(startPoint, endPoint Point, color Color) *Image {
	for x := min(startPoint.X, endPoint.X); x <= max(startPoint.X, endPoint.X); x++ {
		for y := min(startPoint.Y, endPoint.Y); y <= max(startPoint.X, endPoint.Y); y++ {
			if isPointOnLine(Point{x, y}, startPoint, endPoint) {
				image.ImageData.Set(x, y, color)
			}
		}
	}
	return image
}

func (image *Image) StrokeCircle(centerPoint Point, radius int, color Color) *Image {
	for theta := 0.0; theta <= math.Pi*2; theta += 0.017 {
		x := float32(math.Cos(theta))*float32(radius) + float32(centerPoint.X)
		y := float32(math.Sin(theta))*float32(radius) + float32(centerPoint.Y)
		image.ImageData.Set(int(x), int(y), color)
	}
	return image
}

func (image *Image) FillCircle(centerPoint Point, radius int, color Color) *Image {
	diameter := radius
	startPoint := Point{X: centerPoint.X - diameter, Y: centerPoint.Y - diameter}
	endPoint := Point{X: centerPoint.X + diameter, Y: centerPoint.Y + diameter}
	for x := min(startPoint.X, endPoint.X); x <= max(startPoint.X, endPoint.X); x++ {
		for y := min(startPoint.Y, endPoint.Y); y <= max(startPoint.Y, endPoint.Y); y++ {

			if isPointInsideCircle(Point{X: x, Y: y}, centerPoint, radius) {
				image.ImageData.Set(int(x), int(y), color)
			}
		}
	}
	return image
}

func (image Image) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, image.ImageData)
	return nil
}

func CreateImage(width, height int) (p Image) {
	startPoint := image.Point{0, 0}
	endPoint := image.Point{width, height}
	p.ImageData = image.NewRGBA(image.Rectangle{startPoint, endPoint})
	p.Width = width
	p.Height = height
	return p
}
