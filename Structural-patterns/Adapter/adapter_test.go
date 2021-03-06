package Adapter

import "testing"

// 适配器——类对象结构型模式
func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	adapter := PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{},
		Msg:        msg,
	}
	returnedMsg := adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{
		OldPrinter: nil,
		Msg:        msg,
	}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
