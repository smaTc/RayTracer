package graphics

import "os"

import "strconv"

import "fmt"

//CreateCanvas Function
func CreateCanvas(w, h int) Canvas {
	pixels := make([][]Pixel, h)
	for i := range pixels {
		pixels[i] = make([]Pixel, w)
		for j := range pixels[i] {
			pixels[i][j].Color = Color{Red: 0, Green: 0, Blue: 0}
		}
	}
	return Canvas{Pixels: pixels}
}

//WritePixelToCanvas Function
func (c *Canvas) WritePixelToCanvas(x, y int, cl *Color) {
	height := len(c.Pixels) - 1
	width := len(c.Pixels[0]) - 1
	if height-y < 0 || x > width || x < 0 {
		return
	}
	WritePixel(c, x, height-y, cl)
}

//WritePixel Function
func WritePixel(c *Canvas, x, y int, cl *Color) {
	c.Pixels[y][x].Color.Red = cl.Red
	c.Pixels[y][x].Color.Green = cl.Green
	c.Pixels[y][x].Color.Blue = cl.Blue
}

//SaveCanvasAsPPM Function
func SaveCanvasAsPPM(c *Canvas, path, name string) {
	os.Remove(path + name + ".ppm")
	f, err := os.Create(path + name + ".ppm")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("created")

	f.WriteString(canvasToString(c))
	fmt.Println("written")
	f.Sync()
	fmt.Println("synced")
}

func adjustColors(p *Pixel) {
	clamp(&p.Color.Red)
	clamp(&p.Color.Green)
	clamp(&p.Color.Blue)

}

func clamp(v *float64) {
	if *v > 1 {
		*v = 1
	} else if *v < 0 {
		*v = 0
	}
}

func canvasToString(c *Canvas) string {
	w := strconv.Itoa(len(c.Pixels))
	h := strconv.Itoa(len(c.Pixels[0]))
	var str string = ""
	str += "P3\n" + w + " " + h + "\n" + "255\n"
	for i, pixels := range c.Pixels {
		for j, pixel := range pixels {
			adjustColors(&pixel)
			r := strconv.Itoa(int(pixel.Color.Red * 255))
			g := strconv.Itoa(int(pixel.Color.Green * 255))
			b := strconv.Itoa(int(pixel.Color.Blue * 255))
			str += r + " " + g + " " + b
			if j != len(pixels)-1 {
				str += " "
			} else {
				str += "\n"
			}
			fmt.Println("Writing:"+strconv.Itoa(j), strconv.Itoa(i))
		}
	}
	return str
}
