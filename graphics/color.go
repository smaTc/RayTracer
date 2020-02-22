package graphics

//Color struct
type Color struct {
	Red, Green, Blue float64
}

//AddColor Function
func (c *Color) AddColor(c1 *Color) Color {
	return Color{c.Red + c1.Red, c.Green + c1.Green, c.Blue + c1.Blue}
}

//SubColor Function
func (c *Color) SubColor(c1 *Color) Color {
	return Color{c.Red - c1.Red, c.Green - c1.Green, c.Blue - c1.Blue}
}

//MulColor Function
func (c *Color) MulColor(c1 *Color) Color {
	return Color{c.Red * c1.Red, c.Green * c1.Green, c.Blue * c1.Blue}
}

//Scale Function
func (c *Color) Scale(s float64) Color {
	return Color{c.Red * s, c.Green * s, c.Blue * s}
}

//MulColors Function
func MulColors(c1, c2 *Color) Color {
	return c1.MulColor(c2)
}

//AddColors Function
func AddColors(c1, c2 *Color) Color {
	return c1.AddColor(c2)
}

//SubColors Function
func SubColors(c1, c2 *Color) Color {
	return c1.SubColor(c2)

}
