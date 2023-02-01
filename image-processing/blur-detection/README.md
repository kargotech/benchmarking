# benchmarking
This directory contains code which runs various blur detection algorithms against a test suite. 

The test suite can be run by simply running benchmark.go. 
`go run benchmark.go`

The benchmark.go file also contains a training function which can be used to determine the optimal value against a specific test set for the blur threshold of any algorithm.

Dataset is not included in git.

### Dataset
All images were pulled from publically available datasets. 
Testing and training data, required labeled data which was aquired from 
https://paperswithcode.com/dataset/real-blur-dataset.

High and low-res photos were taken from random datasets and filtered for appropriate size requirements. 

Data set directory structure (must be recreated to rerun benchmarking)
```
blur-detection/
    sample-images/
        high-res/
        low-res/
        testing/
            blurry/
            sharp/
        training/
            blurry/
            sharp/
```


### Results 
```
LIBRARY DATASET INFORMATION
TRAINING LIBRARY
	Number of Images: 100
	Average Size: 761 KB
TESTING LIBRARY
	Number of Images: 100
	Average Size: 761 KB
HIGH-RES LIBRARY
	Number of Images: 50
	Average Size: 1929 KB
LOW-RES LIBRARY
	Number of Images: 50
	Average Size: 451 KB

Testing AWS_REKOGNITION
	Blur Detection Success Rate: 84.000000
	Blur Detection Average Speed: 986
	Sharp Detection Success Rate: 80.000000
	Sharp Detection Average Speed: 1128
	Overall Accurracy: 82.000000
	Overall Speed: 1057.100000
	Overall Speed Against HighRes Library: 1845.880000
	Overall Speed Against LowRes Library: 861.400000
Testing LAPLACIAN_VARIANCE
	Blur Detection Success Rate: 86.000000
	Blur Detection Average Speed: 131
	Sharp Detection Success Rate: 76.000000
	Sharp Detection Average Speed: 130
	Overall Accurracy: 81.000000
	Overall Speed: 131.130000
	Overall Speed Against HighRes Library: 310.000000
	Overall Speed Against LowRes Library: 137.000000
```
