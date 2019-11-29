package graphics

//AddColor Function
func (c Color) AddColor(c1 Color) Color {
	return Color{c.Red + c1.Red, c.Green + c1.Green, c.Blue + c1.Blue}
}

//SubColor Function
func (c Color) SubColor(c1 Color) Color {
	return Color{c.Red - c1.Red, c.Green - c1.Green, c.Blue - c1.Blue}
}

//MulColor Function
func (c Color) MulColor(c1 Color) Color {
	return Color{c.Red * c1.Red, c.Green * c1.Green, c.Blue * c1.Blue}
}

//Scale Function
func (c Color) Scale(s float64) Color {
	return Color{c.Red * s, c.Green * s, c.Blue * s}
}
