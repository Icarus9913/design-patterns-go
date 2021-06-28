package Publish_Subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestStdoutPrinter(t *testing.T) {

}

func TestWriter(t *testing.T) {

}

func TestPublisher1(t *testing.T) {
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)

	sub := NewWriterSubscriber(0, nil)

	stdoutPrinter := sub.(*writerSubscriber)
	// 将sub的Writer替换掉,因为Writer接口赋的实际值是指针, fmt.Fprintf会执行这个Write方法
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			fmt.Println(res)
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	sub.Close()
}
