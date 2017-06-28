package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

type imgRGBA [][]float64

func dataToImage(data []byte, imageName string) (image.Image, error) {
	//var histogram imgRGBA
	reader := bytes.NewReader(data)
	//var imageExtension = strings.Split(imageName, ".")[1]
	var img image.Image
	var err error
	/*switch imageExtension {
	case "png":
		img, err = png.Decode(reader)
	case "jpg":
		img, err = jpeg.Decode(reader)
	case "jpeg":
		img, err = jpeg.Decode(reader)
	default:
		img = nil
	}
	*/
	img, err = jpeg.Decode(reader)
	if err != nil {
		return img, err
	}
	return img, err
}

func imageToData(img image.Image, imageName string) ([]byte, error) {
	buf := new(bytes.Buffer)
	//var imageExtension = strings.Split(imageName, ".")[1]
	var err error
	/*switch imageExtension {
	case "png":
		err = png.Encode(buf, img)
	case "jpg":
		err = jpeg.Encode(buf, img, nil)
	case "jpeg":
		err = jpeg.Encode(buf, img, nil)
	default:
		img = nil
	}*/
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), err
}

func imageToPNG(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	var err error
	err = png.Encode(buf, img)
	return buf.Bytes(), err
}

func imageToHistogram(img image.Image) [][]float64 {
	bounds := img.Bounds()

	//generate the histogram
	var histogram [][]float64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			var pixel []float64
			pixel = append(pixel, float64(r), float64(g), float64(b), float64(a))
			histogram = append(histogram, pixel)
		}
	}
	return histogram
}

func Resize(img image.Image) image.Image {
	r := resize.Resize(uint(config.ImgWidth), uint(config.ImgHeigh), img, resize.Lanczos3)
	return r
}