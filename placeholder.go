package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var watermark string = "iVBORw0KGgoAAAANSUhEUgAAAGQAAAAZCAYAAADHXotLAAABhGlDQ1BJQ0MgcHJvZmlsZQAAKJF9kT1Iw0AYht+mSkUqDlaQ4pChOlkQFXHUKhShQqgVWnUwufQPmjQkKS6OgmvBwZ/FqoOLs64OroIg+APi7OCk6CIlfpcUWsR4x3EP733vy913gNCoMM3qGgc03TbTyYSYza2KoVcEMUhTQFRmljEnSSn4jq97BPh+F+dZ/nV/jj41bzEgIBLPMsO0iTeIpzdtg/M+cYSVZJX4nHjMpAsSP3Jd8fiNc9FlgWdGzEx6njhCLBY7WOlgVjI14inimKrplC9kPVY5b3HWKjXWuid/YTivryxzndYwkljEEiSIUFBDGRXYiNOuk2IhTecJH3/U9UvkUshVBiPHAqrQILt+8D/43VurMDnhJYUTQPeL43yMAKFdoFl3nO9jx2meAMFn4Epv+6sNYOaT9Hpbix0B/dvAxXVbU/aAyx1g6MmQTdmVgrSEQgF4P6NvygEDt0Dvmte31jlOH4AM9Sp1AxwcAqNFyl73eXdPZ9/+rWn17wcqznKKY53moAAAAAZiS0dEAP8A/wD/oL2nkwAAAAlwSFlzAAAuIwAALiMBeKU/dgAAANhJREFUaN7tl8kSwyAMQ1GH//9l9dRbCF5E2uno3cIQkBewGcMYY4wxxpi/B4txbuZy839kDm90cKOXCU0s2sGAv6I6o34fr0TwWAw6L74jYwootGMINS7XmYkT9HEaRMKRyLSO0Wg6NOIDCGxNnxDllXhlEB7YH+JMlzOTmY5GZqHgiN3+PHy1nNKJakAqweBmDQr3xE2T8GRQUAggujVEWQBZ6NqqGjC+T7iGzENHFYsuCoKgdrKx0xqPZossqSEn2kIs2lwcyEoG3jzGGGOMMcYYY8zv8AY840AYvp+dGgAAAABJRU5ErkJggg=="
var watermarkAngled string = "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"

func applyWatermark(original image.Image, watermark image.Image, useWatermarkColor bool, watermarkColor *color.RGBA) image.Image {
	originalBounds := original.Bounds()
	watermarkBounds := watermark.Bounds()

	// combine image with watermark
	watermarked := image.NewRGBA(originalBounds)
	for y := 0; y < originalBounds.Dy(); y++ {
		for x := 0; x < originalBounds.Dx(); x++ {
			// get coords for watermark image
			wx := ((x + watermarkBounds.Dx()/2) % watermarkBounds.Dx())
			wy := ((y + watermarkBounds.Dy()/2) % watermarkBounds.Dy())
			wc := color.RGBAModel.Convert(watermark.At(wx, wy)).(color.RGBA)
			// check if watermark if transparent at this location
			if wc.A == 0 {
				// if yes, use original pixel
				c := color.RGBAModel.Convert(original.At(x, y)).(color.RGBA)
				watermarked.SetRGBA(x, y, c)
			} else {
				// if no, use watermark pixel
				if useWatermarkColor {
					watermarked.SetRGBA(x, y, *watermarkColor)
				} else {
					watermarked.SetRGBA(x, y, wc)
				}
			}
		}
	}

	return watermarked
}

func checkFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !(errors.Is(err, os.ErrNotExist) || errors.Is(err, os.ErrPermission))
}

func main() {
	// TODO: allow for output path setting
	// TODO: allow for watermark image overriding?
	constColor := &color.RGBA{255, 0, 0, 255}

	// Prepare for parsing command line args
	trueConst := true

	useConstColor := &trueConst
	tiltWatermark := &trueConst
	// Custom Usage message
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s file\n", os.Args[0])
	}

	// Parse args
	flag.Parse()

	// Check if user provided at least 1 argument
	if len(os.Args) < 2 {
		fmt.Fprintf(flag.CommandLine.Output(), "file not provided\n")
		flag.Usage()
		os.Exit(2)
	}

	// Get filename (last arg)
	filename := os.Args[len(os.Args)-1]

	// Check if file exists
	if !checkFileExists(filename) {
		fmt.Fprintf(flag.CommandLine.Output(), "file %s does not exist or is not readable\n", filename)
		flag.Usage()
		os.Exit(1)
	}

	// read file
	reader, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "could not read %s: %s\n", filename, err.Error())
		flag.Usage()
		os.Exit(1)
	}
	defer reader.Close()

	// decode image from file
	original, formatName, err := image.Decode(reader)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "could not decode %s: %s\n", filename, err.Error())
		flag.Usage()
		os.Exit(1)
	}

	// prepare watermark image
	var readerWatermark io.Reader
	if *tiltWatermark {
		readerWatermark = base64.NewDecoder(base64.StdEncoding, strings.NewReader(watermarkAngled))
	} else {
		readerWatermark = base64.NewDecoder(base64.StdEncoding, strings.NewReader(watermark))
	}
	watermark, _, err := image.Decode(readerWatermark)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "could not decode watermark image: %s\n", err.Error())
		flag.Usage()
		os.Exit(3)
	}

	// apply watermark
	watermarked := applyWatermark(original, watermark, *useConstColor, constColor)

	// save file as watermarked_<old name>
	outputName := "watermarked_" + filepath.Base(filename)
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "could not create output file %s: %s\n", outputName, err.Error())
		flag.Usage()
		os.Exit(1)
	}
	defer outFile.Close()

	if formatName == "png" {
		png.Encode(outFile, watermarked)
	} else if formatName == "jpeg" {
		jpeg.Encode(outFile, watermarked, &jpeg.Options{Quality: 90})
	}
}
