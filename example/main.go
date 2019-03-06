package main

import (
	"github.com/hedarikun/photon"
)

func main() {
	image := photon.CreateImage(400, 400)
	image.FillRect(photon.Point{0, 0}, photon.Point{100, 100}, photon.Color{100, 100, 100, 100})
	image.StrokeRect(photon.Point{200, 200}, photon.Point{300, 300}, 5, photon.Color{100, 100, 200, 255})
	image.StrokeLine(photon.Point{0, 0}, photon.Point{400, 400}, photon.Color{100, 200, 255, 255})
	image.StrokeCircle(photon.Point{200, 200}, 55, photon.Color{255, 255, 0, 255})
	image.FillCircle(photon.Point{100, 100}, 20, photon.Color{243, 13, 21, 255})
	image.Save("something.png")
}
