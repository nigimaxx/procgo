package pkg

import (
	"strings"

	"github.com/fatih/color"
)

var (
	colorIndex = 0
	colors     = []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgWhite,
	}
)

func NewPrefixWriter(prefix string, ch chan []byte) PrefixWriter {
	colorIndex = (colorIndex + 1) % len(colors)
	return PrefixWriter{prefix, colors[colorIndex], ch}
}

type PrefixWriter struct {
	prefix    string
	colorName color.Attribute
	ch        chan []byte
}

func (w PrefixWriter) Write(p []byte) (n int, err error) {
	split := strings.Split(string(p), "\n")
	prefix := color.New(w.colorName).Sprintf("%-20s | ", w.prefix)

	withPrefix := ""

	for _, s := range split {
		if len(s) > 0 {
			withPrefix += prefix + s + "\n"
		}
	}

	w.ch <- []byte(withPrefix)

	return len(p), nil
}
