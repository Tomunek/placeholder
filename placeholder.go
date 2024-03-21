package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

func main() {
	// TODO: parse cmdline args
	constColor := false
	angled := true

	// TODO: load image
	reader, err := os.Open("test.png")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
	defer reader.Close()

	original, formatName, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	// TODO: add watermark
	var readerWatermark io.Reader
	if angled {
		readerWatermark = base64.NewDecoder(base64.StdEncoding, strings.NewReader(watermarkAngled))
	} else {
		readerWatermark = base64.NewDecoder(base64.StdEncoding, strings.NewReader(watermark))
	}

	watermark, _, err := image.Decode(readerWatermark)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	watermarkBounds := watermark.Bounds()

	bounds := original.Bounds()
	watermarked := image.NewRGBA(bounds)

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			c := color.RGBAModel.Convert(original.At(x, y)).(color.RGBA)
			wx := ((x + watermarkBounds.Dx()/2) % watermarkBounds.Dx())
			wy := ((y + watermarkBounds.Dy()/2) % watermarkBounds.Dy())
			wc := color.RGBAModel.Convert(watermark.At(wx, wy)).(color.RGBA)
			if wc.A == 0 {
				// use original pixel
				watermarked.SetRGBA(x, y, c)
			} else {
				// use watermark pixel
				if constColor {
					watermarked.SetRGBA(x, y, color.RGBA{255, 0, 0, 255})
				} else {
					watermarked.SetRGBA(x, y, wc)
				}
			}
		}
	}

	if formatName == "png" {
		outFile, _ := os.Create("./output.png")
		png.Encode(outFile, watermarked)
	} else if formatName == "jpeg" {
		outFile, _ := os.Create("./output.jpg")
		jpeg.Encode(outFile, watermarked, &jpeg.Options{Quality: 90})
	}
}

var watermark string = "iVBORw0KGgoAAAANSUhEUgAAAGQAAAAZCAYAAADHXotLAAABhGlDQ1BJQ0MgcHJvZmlsZQAAKJF9kT1Iw0AYht+mSkUqDlaQ4pChOlkQFXHUKhShQqgVWnUwufQPmjQkKS6OgmvBwZ/FqoOLs64OroIg+APi7OCk6CIlfpcUWsR4x3EP733vy913gNCoMM3qGgc03TbTyYSYza2KoVcEMUhTQFRmljEnSSn4jq97BPh+F+dZ/nV/jj41bzEgIBLPMsO0iTeIpzdtg/M+cYSVZJX4nHjMpAsSP3Jd8fiNc9FlgWdGzEx6njhCLBY7WOlgVjI14inimKrplC9kPVY5b3HWKjXWuid/YTivryxzndYwkljEEiSIUFBDGRXYiNOuk2IhTecJH3/U9UvkUshVBiPHAqrQILt+8D/43VurMDnhJYUTQPeL43yMAKFdoFl3nO9jx2meAMFn4Epv+6sNYOaT9Hpbix0B/dvAxXVbU/aAyx1g6MmQTdmVgrSEQgF4P6NvygEDt0Dvmte31jlOH4AM9Sp1AxwcAqNFyl73eXdPZ9/+rWn17wcqznKKY53moAAAAAZiS0dEAP8A/wD/oL2nkwAAAAlwSFlzAAAuIwAALiMBeKU/dgAAANhJREFUaN7tl8kSwyAMQ1GH//9l9dRbCF5E2uno3cIQkBewGcMYY4wxxpi/B4txbuZy839kDm90cKOXCU0s2sGAv6I6o34fr0TwWAw6L74jYwootGMINS7XmYkT9HEaRMKRyLSO0Wg6NOIDCGxNnxDllXhlEB7YH+JMlzOTmY5GZqHgiN3+PHy1nNKJakAqweBmDQr3xE2T8GRQUAggujVEWQBZ6NqqGjC+T7iGzENHFYsuCoKgdrKx0xqPZossqSEn2kIs2lwcyEoG3jzGGGOMMcYYY8zv8AY840AYvp+dGgAAAABJRU5ErkJggg=="
var watermarkAngled string = "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"
