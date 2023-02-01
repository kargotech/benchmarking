package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	bluralgorithms "github.com/kargotech/benchmarking/image-processing/blur-detection/blur-algorithms"
)

var cur_directory, _ = os.Getwd()
var sample_images = filepath.Join(cur_directory, "sample-images")
var training_library = filepath.Join(sample_images, "training")
var testing_library = filepath.Join(sample_images, "testing")
var highres_library = filepath.Join(sample_images, "high-res")
var lowres_library = filepath.Join(sample_images, "low-res")
var blurry_library = "blurry"
var sharp_library = "sharp"

func main() {
	printLibraryInformation()
	algs := bluralgorithms.GetAlgorithms()
	for _, v := range algs {
		// train_algorithm(v)
		test_algorithm(v)
	}
}

// Main diagnostic function
func test_algorithm(g bluralgorithms.BlurAlgorithm) {
	fmt.Println("Testing " + g.Name())
	threshold := g.GetThreshold()

	// Get all blurry photos from the training library
	d_blurry, _ := os.ReadDir(filepath.Join(training_library, blurry_library))
	total_blurry := len(d_blurry)
	total_blur_detected := 0
	total_time_elapsed := float64(0)

	time_elapsed := float64(0)
	// Calculate the blur value of each photo and determine if the value is lower than the threshold
	for _, v := range d_blurry {
		filepath := filepath.Join(training_library, blurry_library, v.Name())
		value, time := g.GetValue(filepath)
		if value < threshold {
			total_blur_detected++
		}
		time_elapsed += time
	}
	fmt.Printf("\tBlur Detection Success Rate: %f\n", float64(total_blur_detected)/float64(total_blurry)*100)
	fmt.Printf("\tBlur Detection Average Speed: %d\n", int(time_elapsed/float64(total_blurry)))

	// Get all sharp photos from the training library
	d_sharp, _ := os.ReadDir(filepath.Join(training_library, sharp_library))
	total_sharp := len(d_sharp)
	total_sharp_detected := 0
	total_time_elapsed += time_elapsed
	time_elapsed = float64(0)

	// Calculate the blur value of each sharp photo and determine if the value is higher than the threshold
	for _, v := range d_sharp {
		filepath := filepath.Join(training_library, sharp_library, v.Name())
		value, time := g.GetValue(filepath)
		if value > threshold {
			total_sharp_detected++
		}
		time_elapsed += time
	}
	total_time_elapsed += time_elapsed
	fmt.Printf("\tSharp Detection Success Rate: %f\n", float64(total_sharp_detected)/float64(total_sharp)*100)
	fmt.Printf("\tSharp Detection Average Speed: %d\n", int(time_elapsed/float64(total_sharp)))

	fmt.Printf("\tOverall Accurracy: %f\n", float64(total_sharp_detected+total_blur_detected)/float64(total_sharp+total_blurry)*100)
	fmt.Printf("\tOverall Speed: %f\n", float64(total_time_elapsed)/float64(total_sharp+total_blurry))

	fmt.Printf("\tOverall Speed Against HighRes Library: %f\n", runSpeedTest(g, highres_library))
	fmt.Printf("\tOverall Speed Against LowRes Library: %f\n", runSpeedTest(g, lowres_library))

}

// Gets average speed of blur detection algorithm against a specific library
func runSpeedTest(g bluralgorithms.BlurAlgorithm, library string) float64 {
	library_dir, _ := os.ReadDir(library)
	library_size := len(library_dir)

	time_elapsed := float64(0)
	for _, v := range library_dir {
		filepath := filepath.Join(library, v.Name())
		_, time := g.GetValue(filepath)
		time_elapsed += time
	}
	return time_elapsed / float64(library_size)
}

func printLibraryInformation() {
	fmt.Println("LIBRARY DATASET INFORMATION")
	fmt.Println("TRAINING LIBRARY")
	analyze_labeled_library(training_library)
	fmt.Println("TESTING LIBRARY")
	analyze_labeled_library(training_library)
	fmt.Println("HIGH-RES LIBRARY")
	analyze_library(highres_library)
	fmt.Println("LOW-RES LIBRARY")
	analyze_library(lowres_library)
}

func analyze_library(library_path string) {
	d, e := os.ReadDir(library_path)
	if e != nil {
		panic(e)
	}
	analyze_library_helper(d)
}

func analyze_labeled_library(library_path string) {
	d_blurry, e := os.ReadDir(filepath.Join(library_path, blurry_library))
	d_sharp, e := os.ReadDir(filepath.Join(library_path, sharp_library))
	d := append(d_blurry, d_sharp...)
	if e != nil {
		panic(e)
	}
	analyze_library_helper(d)
}

func analyze_library_helper(d []fs.DirEntry) {
	var total_size int64 = 0
	for _, v := range d {
		info, _ := v.Info()
		total_size += info.Size()
	}

	fmt.Printf("\tNumber of Images: %d\n", len(d))
	fmt.Printf("\tAverage Size: %d KB\n", total_size/(1000*int64(len(d))))
}

// Determines Optimal Threshold Value for a blur algorithm
func train_algorithm(g bluralgorithms.BlurAlgorithm) {
	fmt.Println("Training " + g.Name())
	d_blurry, _ := os.ReadDir(filepath.Join(training_library, blurry_library))
	var blurry_values []float64
	for _, v := range d_blurry {
		filepath := filepath.Join(training_library, blurry_library, v.Name())
		value, _ := g.GetValue(filepath)
		blurry_values = append(blurry_values, value)
	}

	d_sharp, _ := os.ReadDir(filepath.Join(training_library, sharp_library))
	var sharp_values []float64
	for _, v := range d_sharp {
		filepath := filepath.Join(training_library, sharp_library, v.Name())
		value, _ := g.GetValue(filepath)
		sharp_values = append(sharp_values, value)
	}

	optimal_threshold := getScores(blurry_values, sharp_values, g)
	fmt.Printf("OPTIMAL THRESHOLD: %f\n", optimal_threshold)

}

func getScores(blurry_values []float64, sharp_values []float64, g bluralgorithms.BlurAlgorithm) float64 {
	min, max, increment := g.GetThresholdRange()
	var max_score float64
	var score float64
	var optimal_threshold float64
	for i := min; i <= max; i += increment {
		score = 0
		for _, v := range blurry_values {
			if v < i {
				score += 1
			}
		}
		for _, v := range sharp_values {
			if v > i {
				score += 1
			}
		}
		if i == min || score > max_score {
			max_score = score
			optimal_threshold = i
		}

	}
	return optimal_threshold
}
