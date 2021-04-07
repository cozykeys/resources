package kbutil

import (
	"strings"
	//"xml"
)

func (kb *Keyboard) ToSvg(tags []string) (string, error) {
	// TODO: Figure out how to use the xml library instead of hard coding all
	// of the element/attribute syntax
	lines := []string{
		//string('\ufeff') + "<?xml version=\"1.0\" encoding=\"utf-8\"?>",
		"<?xml version=\"1.0\" encoding=\"utf-8\"?>",
		"<svg",
		"    width=\"500mm\"",
		"    height=\"200mm\"",
		"    viewBox=\"0 0 500 200\" xmlns=\"http://www.w3.org/2000/svg\">",
		"</svg>",
	}

	return strings.Join(lines, "\n"), nil
}
