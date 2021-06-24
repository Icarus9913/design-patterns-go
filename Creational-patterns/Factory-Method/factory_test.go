package Factory_Method

import (
	"strings"
	"testing"
)

// 工厂方法
func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMehod(Cash)
	if nil != err {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMehod(t *testing.T) {
	payment, err := GetPaymentMehod(DebitCard)
	if nil != err {
		t.Error("A payment method of type 'DebitCard' must exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	payment, err := GetPaymentMehod(20)
	if nil == err {
		t.Error("A payment method with ID 20 must return an error")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}
