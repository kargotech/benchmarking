package bluralgorithms

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type AWS_REKOGNITION struct {
}

func (r AWS_REKOGNITION) GetThreshold() float64 {
	return 40
}

func (r AWS_REKOGNITION) GetValue(filePath string) (float64, float64) {
	return runRekog(filePath)
}

func (r AWS_REKOGNITION) Name() string {
	return "AWS_REKOGNITION"
}

func (r AWS_REKOGNITION) GetThresholdRange() (float64, float64, float64) {
	return 0, 100, 5
}

func runRekog(filePath string) (float64, float64) {
	region := "ap-southeast-1"
	svc := rekognition.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	})))
	file, _ := ioutil.ReadFile(filePath)

	GENERAL_LABELS := "GENERAL_LABELS"
	IMAGE_PROPERTIES := "IMAGE_PROPERTIES"

	input := &rekognition.DetectLabelsInput{
		Features: []*string{&GENERAL_LABELS, &IMAGE_PROPERTIES},
		Image: &rekognition.Image{
			Bytes: file,
		},
		MaxLabels:     aws.Int64(1),
		MinConfidence: aws.Float64(100.000000),
	}
	start := time.Now()
	result, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Println(err)
	}
	sharpness := *result.ImageProperties.Quality.Sharpness
	elapsed := time.Since(start)
	return sharpness, float64(elapsed.Milliseconds())
}
