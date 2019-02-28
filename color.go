package photon

// Color is photon color object that uses RGBA values for each pixel in image
type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

// SetColor is a function that is used to set the color value for photon color object
func (c *Color) SetColor(red, green, blue, alpha uint8) {
	c.Red = red
	c.Green = green
	c.Blue = blue
	c.Alpha = alpha
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.Red)
	r |= r << 8
	g = uint32(c.Green)
	g |= g << 8
	b = uint32(c.Blue)
	b |= b << 8
	a = uint32(c.Alpha)
	a |= a << 8
	return
}
