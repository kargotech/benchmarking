package bluralgorithms

import (
	"image"
	"image/jpeg"
	"os"
	"time"

	"gocv.io/x/gocv"
)

type LAPLACIAN_VARIANCE_ALG struct {
}

func (r LAPLACIAN_VARIANCE_ALG) GetThreshold() float64 {
	return 140
}

func (r LAPLACIAN_VARIANCE_ALG) GetThresholdRange() (float64, float64, float64) {
	return 0, 500, 10
}

func (r LAPLACIAN_VARIANCE_ALG) GetValue(filePath string) (float64, float64) {
	f, _ := os.Open(filePath)
	defer f.Close()
	image, _ := jpeg.Decode(f)
	start := time.Now()
	variance := DetectBlur(image)
	elapsed := time.Since(start)
	return variance, float64(elapsed.Milliseconds())
}

func (r LAPLACIAN_VARIANCE_ALG) Name() string {
	return "LAPLACIAN_VARIANCE"
}

func DetectBlur(Img image.Image) float64 {

	// struct initialization
	tempMat := gocv.NewMat()
	defer tempMat.Close()

	imgCV, _ := gocv.ImageToMatRGB(Img)
	defer imgCV.Close()

	gray := gocv.NewMat()
	defer gray.Close()

	// Change matrix to grayscale
	gocv.CvtColor(imgCV, &gray, gocv.ColorBGRToGray)

	blur := gocv.NewMat()
	defer blur.Close()

	// Apply Laplacian filter for rapid changes of neighboring pixels
	gocv.Laplacian(gray, &blur, gocv.MatTypeCV64F, 1, 1, 0, gocv.BorderDefault)

	mean := gocv.NewMat()
	defer mean.Close()

	stdDev := gocv.NewMat()
	defer stdDev.Close()

	gocv.MeanStdDev(blur, &mean, &stdDev)

	stdDevFloat, _ := stdDev.DataPtrFloat64()

	// Variance = stdev^2
	variance := stdDevFloat[0] * stdDevFloat[0]
	return variance
}
