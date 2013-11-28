package newrelic_go_api

import (
	"fmt"
	"testing"
	"time"
)

func TestTransactionSend(t *testing.T) {

	Init(NEWRELIC_LICENSE, TEST_APP_NAME)
	time.Sleep(10 * time.Second)
	SetupEmbeddedCollectorClient()

	trId := StartWebTransaction()
	if trId == 0 {
		t.Errorf("StartWebTransaction failed. Return 0 transaction ID")
	}

	if result := NameWebTransaction(trId, "Unit test transaction"); result != 0 {
		t.Errorf("NameWebTransaction failed. Return code: %d", result)
	}

	time.Sleep(10 * time.Millisecond)
	if result := EndWebTransaction(trId); result != 0 {
		t.Errorf("NameWebTransaction failed. Return code: %d", result)
	}
	fmt.Printf("Send transaction with ID: %d and wait for 60 seconds\n", trId)
	time.Sleep(70 * time.Second)

	if result := RequestShutdown("Reason: transaction test"); result != 0 {
		t.Errorf("Shutdown test failed. Return: %d.", result)
	}
	time.Sleep(10 * time.Second)
}

/*
func TestStartWebTransaction(t *testing.T) {
	if result := StartWebTransaction(); result == 0 {
		t.Errorf("StartWebTransaction failed. Return 0 transaction ID")
	}
}

func TestNameWebTransaction(t *testing.T) {
	trId := StartWebTransaction()
	if result := NameWebTransaction(trId, "Unit test transaction"); result != 0 {
		t.Errorf("NameWebTransaction failed. Return code: %d", result)
	}
}
*/
