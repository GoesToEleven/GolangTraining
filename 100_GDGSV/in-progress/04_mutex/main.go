package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

type pixel struct {
	r, g, b, a uint32
}

func main() {
	start := time.Now()
	dir := "../photos/"
	files, _ := ioutil.ReadDir(dir)
	wg.Add(len(files))
	fmt.Println("FILES TO PROCESS:", len(files))
	images := getImages(dir)

	// range over the [] holding the []pixel - eg, give me each img
	//     range over the []pixel hold the pixels - eg, give me each pixel
	for i, img := range images {
		for j, pixel := range img {
			fmt.Println("Image", i, "\t pixel", j, "\t r g b a:", pixel)
			if j == 10 {
				break
			}
		}
	}
	wg.Wait()
	fmt.Println("Holy cow:", counter)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getImages(dir string) [][]pixel {

	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		go (func() {
			img := loadImage(path)
			pixels := getPixels(img)
			mutex.Lock()
			images = append(images, pixels)
			mutex.Unlock()
			wg.Done()
		})()
		return nil
	})

	return images
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func getPixels(img image.Image) []pixel {

	bounds := img.Bounds()
	fmt.Println(bounds.Dx(), " x ", bounds.Dy()) // debugging
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
		counter++
	}

	return pixels
}
