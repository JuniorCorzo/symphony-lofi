package art

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strings"
)

// ConvertImageToANSI fetches an image from URL and converts it to a 24-bit ANSI block string.
func ConvertImageToANSI(url string, width, height int) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}

	return RenderImageToANSI(img, width, height), nil
}

// RenderImageToANSI converts an image.Image into a 24-bit ANSI truecolor half-block string.
func RenderImageToANSI(src image.Image, targetWidth, targetHeight int) string {
	bounds := src.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	if srcWidth == 0 || srcHeight == 0 {
		return ""
	}

	var result strings.Builder
	totalRows := targetHeight * 2

	for y := 0; y < totalRows; y += 2 {
		srcY1 := bounds.Min.Y + (y * srcHeight / totalRows)
		srcY2 := bounds.Min.Y + ((y + 1) * srcHeight / totalRows)

		for x := range targetWidth {
			srcX := bounds.Min.X + (x * srcWidth / targetWidth)

			topColor := src.At(srcX, srcY1)
			bottomColor := src.At(srcX, srcY2)

			tr, tg, tb, _ := topColor.RGBA()
			br, bg, bb, _ := bottomColor.RGBA()

			tr8, tg8, tb8 := uint8(tr>>8), uint8(tg>>8), uint8(tb>>8)
			br8, bg8, bb8 := uint8(br>>8), uint8(bg>>8), uint8(bb>>8)

			fmt.Fprintf(&result, "\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm▀", tr8, tg8, tb8, br8, bg8, bb8)
		}
		result.WriteString("\x1b[0m\n")
	}

	return result.String()
}
