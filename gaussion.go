package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

// gaussion 二维高斯函数
func gaussion(x, y int, omiga float64) float64 {
	return (1.0 / (2.0 * math.Pi * omiga * omiga)) * math.Pow(math.E, ((-1.0)*(float64(x*x+y*y)/(2.0*omiga*omiga))))
}

// gaussioniMatrix 高斯概率矩阵
func gaussioniMatrix(l int, omiga float64) [][]float64 {
	length := 2*l + 1
	matrics := make([][]float64, length)
	sum := float64(0)
	for i := 0; i < length; i++ {
		arr := make([]float64, length)
		for j := 0; j < length; j++ {
			arr[j] = gaussion(i-l, j-l, omiga)
			sum += arr[j]
		}
		matrics[i] = arr
	}
	for i := range matrics {
		cur := matrics[i]
		for j := range cur {
			cur[j] = cur[j] / sum
		}
		matrics[i] = cur
	}
	return matrics
}

// blur 高斯模糊
func blur(srcImg, destImg string, matrics [][]float64, num int) {
	srcFile, err := os.Open(srcImg)
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destImg)
	if err != nil {
		fmt.Printf("open file err: %v", err)
		return
	}
	defer destFile.Close()

	img, err := jpeg.Decode(srcFile)
	if err != nil {
		fmt.Printf("jpeg decode err: %v", err)
		return
	}
	xW, yH := img.Bounds().Dx(), img.Bounds().Dy()

	jpg := image.NewRGBA64(image.Rect(0, 0, xW-num, yH-num))

	for i := 0; i < xW; i++ {
		for j := 0; j < yH; j++ {
			var newColor color.RGBA64
			var sumR, sumG, sumB, sumA uint16
			for p := ((-1) * num); p <= num; p++ {
				for q := ((-1) * num); q <= num; q++ {
					x, y := i+p, j+q
					if x < 0 {
						x = 0
					} else if x > xW+num {
						x = xW - 1
					}
					if y < 0 {
						y = 0
					} else if y > yH+num {
						y = yH - 1
					}

					R, G, B, A := img.At(x, y).RGBA()
					sumR += uint16(matrics[p+num][q+num] * float64(R))
					sumG += uint16(matrics[p+num][q+num] * float64(G))
					sumB += uint16(matrics[p+num][q+num] * float64(B))
					sumA += uint16(matrics[p+num][q+num] * float64(A))
				}
			}
			newColor.R = sumR
			newColor.G = sumG
			newColor.B = sumB
			newColor.A = sumA
			jpg.SetRGBA64(i, j, newColor)
		}
	}
	// 画图
	jpeg.Encode(destFile, jpg, nil)
	// png.Encode(destFile, jpg)

}

func main() {
	src := "./test.jpg"
	des := "./test_result.jpg"
	num := 5
	omiga := 5.0
	guassionMatrics := gaussioniMatrix(num, omiga)
	blur(src, des, guassionMatrics, num)
}
