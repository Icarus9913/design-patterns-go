package Strategy

import (
	"flag"
	"io"
)

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}
func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

var output = flag.String("output", "image", "The output to use between 'console' and 'image' file")
