package pkg

import (
	"strings"

	"github.com/fatih/color"
)

type PrefixWriter struct {
	prefix    string
	colorName color.Attribute
	ch        chan []byte
}

func (w PrefixWriter) Write(p []byte) (n int, err error) {
	split := strings.Split(string(p), "\n")
	prefix := color.New(w.colorName).Sprintf("%-20s | ", w.prefix)

	withPrefix := prefix + strings.Join(split, "\n"+prefix) + "\n"
	w.ch <- []byte(withPrefix)

	return len(p), nil
}
