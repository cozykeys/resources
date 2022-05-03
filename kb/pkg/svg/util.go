package svg

import (
	"errors"
	"fmt"
	"image/color"
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
