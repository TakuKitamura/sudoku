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
		// handle error
		fmt.Println(err)
	}
}

// HLine draws a horizontal line
func HLine(img *image.RGBA, x1, y, x2 int, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(img *image.RGBA, x, y1, y2 int, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// 画像を単色に染める
func fillRect(img *image.RGBA, col color.Color) {
	// 矩形を取得
	rect := img.Rect

	// 全部埋める
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

	// RectからRGBAを作る(ゼロ値なので黒なはず)
	img := image.NewRGBA(image.Rect(x, y, length, length))
	// 赤色に染める(透過なし)
	fillRect(img, white)
	// Rect(img, size, length-size, size, length-size, black)
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
		HLine(img, size, i, (length - (size)), black)
		VLine(img, i, size, (length - (size)), black)
		for k := 0; k < 9; k++ {
			if j < 9 {
				// fmt.Println(i, j, k, grid[j][k], ((length-(size)*2)/9)*(k+1), i+size)
				if grid[j][k] > 0 {
					num := fmt.Sprint(grid[j][k])
					addLabel(img, (size*2)+(100*k), i+size*4, num)
				}

			}
		}

		if j%3 == 0 {
			HLine(img, size, i+1, (length - (size)), black)
			VLine(img, i+1, size, (length - (size)), black)
			HLine(img, size, i-1, (length - (size)), black)
			VLine(img, i-1, size, (length - (size)), black)
		}
		j++
	}

	// 出力用ファイル作成(エラー処理は略)
	// file, _ := os.Create("sample.jpg")
	// defer file.Close()

	// // JPEGで出力(100%品質)
	// if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
	// 	panic(err)
	// }
}
