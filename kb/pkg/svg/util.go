package svg

import (
	"errors"
	"fmt"
	"image/color"
	"kb/pkg/models"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// reduceSpaces replaces any occurence of 2 or more adjacent spaces with a
// single space.
func reduceSpaces(s string) string {
	return regexp.MustCompile(`  +`).ReplaceAllString(s, " ")
}

var errInvalidFormat = errors.New("invalid format")

func parseColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		default:
			err = errInvalidFormat
			return 0
		}
	}

	c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
	c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
	c.B = hexToByte(s[5])<<4 + hexToByte(s[6])

	return
}

func formatByteHex(c uint8) string {
	hex := strconv.FormatUint(uint64(c), 16)
	if len(hex) < 2 {
		hex = "0" + hex
	}
	return hex
}

func formatColorHex(c color.RGBA) string {
	return "#" + formatByteHex(c.R) + formatByteHex(c.G) + formatByteHex(c.B)
}

func float64OrDefault(f, def float64) float64 {
	if f == 0.0 {
		return def
	}
	return f
}

func stringOrDefault(s, def string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	return s
}

type StyleMap map[string]string

func (m StyleMap) String() string {
	tokens := []string{}
	for k, v := range m {
		tokens = append(tokens, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(tokens, ";")
}

func cssStyleString(m map[string]string) string {
	tokens := []string{}
	for k, v := range m {
		tokens = append(tokens, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(tokens, ";")
}

func getOffsetInStack(stack *models.Stack, element models.KeyboardElement) float64 {
	if stack.Orientation == models.StackOrientationVertical {
		log.Print("getOffsetInStack() - vertical")

		stackHeight := 0.0
		dy := math.MaxFloat64

		for _, child := range stack.GetChildren() {
			childHeight := getTotalHeight(child)
			if child == element {
				dy = stackHeight + childHeight/2
			}
			stackHeight += childHeight
		}

		log.Print(fmt.Sprintf("Stack height = %f", stackHeight))

		return -(stackHeight/2 - dy)
	}

	if stack.Orientation == models.StackOrientationHorizontal {
		log.Print("getOffsetInStack() - horizontal")

		stackWidth := 0.0
		dx := math.MaxFloat64

		for _, child := range stack.GetChildren() {
			childWidth := getTotalWidth(child)
			if child == element {
				dx = stackWidth + childWidth/2
			}
			stackWidth += childWidth
		}

		log.Print(fmt.Sprintf("Stack width = %f", stackWidth))

		return -(stackWidth/2 - dx)
	}

	panic("unknown stack orientation")
}

func getTotalWidth(e models.KeyboardElement) float64 {
	return e.GetWidth() + e.GetMargin()*2
}

func getTotalHeight(e models.KeyboardElement) float64 {
	return e.GetHeight() + e.GetMargin()*2
}

func getMinX(e models.KeyboardElement) float64 {
	return (-getTotalWidth(e) / 2) + e.GetXOffset()
}

func getMaxX(e models.KeyboardElement) float64 {
	return (getTotalWidth(e) / 2) + e.GetXOffset()
}

func getMinY(e models.KeyboardElement) float64 {
	return (-getTotalHeight(e) / 2) + e.GetYOffset()
}

func getMaxY(e models.KeyboardElement) float64 {
	return (getTotalHeight(e) / 2) + e.GetYOffset()
}
