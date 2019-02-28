package photon

import (
	"image"
	"image/png"
	"os"
)

type Image struct {
	ImageData *image.RGBA
	Width     int
	Height    int
}

func (image *Image) FillRect(startPoint, endPoint Point, color Color) *Image {
	// Todo: Check if it can really be drawn in the space of the image
	for x := min(startPoint.X, endPoint.X); x <= max(startPoint.X, endPoint.Y); x++ {
		for y := min(startPoint.Y, endPoint.Y); y <= max(startPoint.X, endPoint.Y); y++ {
			image.ImageData.Set(x, y, color)
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
