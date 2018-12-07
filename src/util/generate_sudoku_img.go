package util

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
)

func addLabel(img *image.RGBA, x, y int, label string) {
	fontBytes, err := ioutil.ReadFile("./HonyaJi-Re.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	c := freetype.NewContext()
	c.SetDst(img)
	c.SetFont(f)
	c.SetFontSize(90.0)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	pt := freetype.Pt(x, y)
	if _, err := c.DrawString(label, pt); err != nil {
		fmt.Println(err)
	}
}

func hLine(img *image.RGBA, x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func vLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

func main() {
	x := 0
	y := 0
	size := 20
	length := 47 * size
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	img := image.NewRGBA(image.Rect(x, y, length, length))
	fillRect(img, white)
	grid := [9][9]uint8{
		{0, 6, 1, 0, 0, 7, 0, 0, 3},
		{0, 9, 2, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 5, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 4},
		{5, 0, 0, 0, 0, 8, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 6, 0, 8, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	j := 0
	for i := size; i < length; i += (length - (size)*2) / 9 {
		hLine(img, size, i, (length - (size)), black)
		vLine(img, i, size, (length - (size)), black)
		for k := 0; k < 9; k++ {
			if j < 9 {
				if grid[j][k] > 0 {
					num := fmt.Sprint(grid[j][k])
					addLabel(img, (size*2)+(100*k), i+size*4, num)
				}
			}
		}
		if j%3 == 0 {
			hLine(img, size, i+1, (length - (size)), black)
			vLine(img, i+1, size, (length - (size)), black)
			hLine(img, size, i-1, (length - (size)), black)
			vLine(img, i-1, size, (length - (size)), black)
		}
		j++
	}
	// file, _ := os.Create("sample.jpg")
	// defer file.Close()

	// if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
	// 	panic(err)
	// }
}
