package gtc

import (
	"bytes"
	"strings"
)

// Useful for comparing logs to expected results.
type WriteCompare struct {
	*bytes.Buffer
	Replacements map[string]string
}

func NewWriteCompare() WriteCompare {
	return WriteCompare{
		new(bytes.Buffer),
		make(map[string]string),
	}
}

func (wc *WriteCompare) String() string {
	str := wc.Buffer.String()
	for original, replace_with := range wc.Replacements {
		str = strings.Replace(str, original, replace_with, -1)
	}
	return str
}
