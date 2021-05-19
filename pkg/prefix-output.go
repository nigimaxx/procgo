package pkg

import (
	"io"
	"os"
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

func NewPrefixOutput(prefix string) PrefixOutput {
	colorIndex = (colorIndex + 1) % len(colors)
	return PrefixOutput{
		Stdout: prefixWriter{prefix, colors[colorIndex], os.Stdout},
		Stderr: prefixWriter{prefix, colors[colorIndex], os.Stderr},
	}
}

type PrefixOutput struct {
	Stdout prefixWriter
	Stderr prefixWriter
}

type prefixWriter struct {
	prefix    string
	colorName color.Attribute
	writer    io.Writer
}

func (w prefixWriter) Write(p []byte) (n int, err error) {
	split := strings.Split(string(p), "\n")
	prefix := color.New(w.colorName).Sprintf("%-20s | ", w.prefix)

	withPrefix := ""

	for _, s := range split {
		if len(s) > 0 {
			withPrefix += prefix + s + "\n"
		}
	}

	if _, err := w.writer.Write([]byte(withPrefix)); err != nil {
		return 0, err
	}

	return len(p), nil
}
