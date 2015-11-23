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

// at terminal:
// go run -race main.go

type pixel struct {
	r, g, b, a uint32
}

var counter int

var wg sync.WaitGroup
var c = make(chan []pixel)

func main() {
	start := time.Now()
	dir := "../photos/"
	files, _ := ioutil.ReadDir(dir)
	wg.Add(len(files))
	fmt.Println("FILES TO PROCESS:", len(files))
	getImages(dir)
	images := puller()

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
	fmt.Println("LENGTH:", len(images))
	fmt.Println("Holy cow:", counter)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func getImages(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		go (func() {
			img := loadImage(path)
			c <- getPixels(img)
		})()
		return nil
	})
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

func puller() [][]pixel {
	var images [][]pixel
	for {
		i, more := <-c
		if more {
			images = append(images, i)
			wg.Done()
		} else {
			close(c)
			return images
		}
	}
}
