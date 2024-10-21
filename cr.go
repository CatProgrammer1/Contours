package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func imageToBlackAndWhite(path string) []color.RGBA {
	var colors []color.RGBA

	file, err := os.Open(path)
	handle(err)

	img, err := png.Decode(file)
	handle(err)

	grayImg := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			grayImg.Set(x, y, grayColor)
		}
	}

	edgeImg := image.NewGray(grayImg.Bounds())
	kernelX := [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	kernelY := [3][3]int{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}

	for x := 1; x < grayImg.Bounds().Max.X-1; x++ {
		for y := 1; y < grayImg.Bounds().Max.Y-1; y++ {
			var gx, gy int
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					grayColor := grayImg.GrayAt(x+i, y+j).Y
					gx += int(grayColor) * kernelX[i+1][j+1]
					gy += int(grayColor) * kernelY[i+1][j+1]
				}
			}
			magnitude := uint8((gx*gx + gy*gy) >> 8)
			edgeImg.SetGray(x, y, color.Gray{magnitude})
		}
	}

	threshold := uint8(30)
	for x := 0; x < edgeImg.Bounds().Max.X; x++ {
		for y := 0; y < edgeImg.Bounds().Max.Y; y++ {
			if edgeImg.GrayAt(x, y).Y > threshold {
				colors = append(colors, color.RGBA{
					R: 255,
					G: 255,
					B: 255,
				})
			} else {

			}
		}
	}

	return colors
}
