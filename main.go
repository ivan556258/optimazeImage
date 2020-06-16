package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

var digits = ""

func main() {
	readDir("/home/z1/Downloads/webstrot.com/html/jbdesk/main_version/main_pages/img/") // here need absolutly patch
}

func readDir(path string) {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() == true {
			newPatch := path + f.Name() + "/"
			fmt.Println(newPatch)
			readDir(newPatch)
		} else {
			fullOatchImg := path + f.Name()
			ConvertToJpeg(fullOatchImg, f.Name())
			ConvertToPng(fullOatchImg, f.Name())

		}
	}
	//return files
}

func ConvertToPng(filename, name string) (string, error) {
	pngImageFile, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer pngImageFile.Close()

	pngSource, err := png.Decode(pngImageFile)

	if err != nil {
		return "", err
	}

	jpegImage := image.NewRGBA(pngSource.Bounds())

	draw.Draw(jpegImage, jpegImage.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(jpegImage, jpegImage.Bounds(), pngSource, pngSource.Bounds().Min, draw.Over)
	fmt.Println(filename)
	outfile := fmt.Sprintf("%s", filename)
	jpegImageFile, err := os.Create(outfile)

	if err != nil {
		return "", err
	}

	defer jpegImageFile.Close()

	var options jpeg.Options
	options.Quality = 50

	err = jpeg.Encode(jpegImageFile, jpegImage, &options)

	if err != nil {
		fmt.Printf("JPEG Encoding Error: %v\n", err)
		os.Exit(1)
	}

	return outfile, nil
}

func ConvertToJpeg(filename, name string) (string, error) {
	jpgImageFile, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer jpgImageFile.Close()

	jpegSource, err := jpeg.Decode(jpgImageFile)

	if err != nil {
		return "", err
	}

	jpegImage := image.NewRGBA(jpegSource.Bounds())

	draw.Draw(jpegImage, jpegImage.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(jpegImage, jpegImage.Bounds(), jpegSource, jpegSource.Bounds().Min, draw.Over)
	fmt.Println(filename)
	outfile := fmt.Sprintf("%s", filename)
	jpegImageFile, err := os.Create(outfile)

	if err != nil {
		return "", err
	}

	defer jpegImageFile.Close()

	var options jpeg.Options
	options.Quality = 50

	err = jpeg.Encode(jpegImageFile, jpegImage, &options)

	if err != nil {
		fmt.Printf("JPEG Encoding Error: %v\n", err)
		os.Exit(1)
	}

	return outfile, nil
}
