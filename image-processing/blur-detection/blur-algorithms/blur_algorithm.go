package bluralgorithms

type BlurAlgorithm interface {
	GetThreshold() float64
	GetValue(filepath string) (float64, float64)
	Name() string
	GetThresholdRange() (float64, float64, float64)
}

func GetAlgorithms() [2]BlurAlgorithm {
	arr := [2]BlurAlgorithm{AWS_REKOGNITION{}, LAPLACIAN_VARIANCE_ALG{}}
	return arr
}
