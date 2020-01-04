// +build example
package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	window := gocv.NewWindow("Hello")
	window2 := gocv.NewWindow("Rotated")
	img := gocv.NewMat()
	rotImg := gocv.NewMat()
	center := image.Point{240, 320}
	for i := 0; ; i++ {
		webcam.Read(&img)
		rotMat := gocv.GetRotationMatrix2D(center, float64(i), 1)
		gocv.WarpAffine(img, &rotImg, rotMat, image.Point{320 * 2, 240 * 2})
		window.IMShow(img)
		window2.IMShow(rotImg)
		window.WaitKey(1)
	}
}
