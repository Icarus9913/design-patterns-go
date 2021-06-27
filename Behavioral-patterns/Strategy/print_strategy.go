package Strategy

import (
	"fmt"
	"io"
	"os"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
