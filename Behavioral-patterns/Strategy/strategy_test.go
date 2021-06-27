package Strategy

import (
	"flag"
	"log"
	"os"
	"testing"
)

// 策略——对象行为型模式
func TestPrint(t *testing.T) {
	flag.Parse()

	activeStrategy, err := NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case IMAGE_STRATEGY:
		//w, err := os.Create("/tmp/image.jpg")
		w, err := os.Create("E:\\image.jpg")
		if err != nil {
			log.Fatal("Error opening image",err)
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}

