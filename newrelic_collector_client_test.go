package newrelic_go_api

import (
	"fmt"
	"testing"
)

const (
	TEST_APP_NAME    = "TestApp"
	NEWRELIC_LICENSE = "7bceac019c7dcafae1ef95be3e3a3ff8866de246"
)

func TestInit(t *testing.T) {
	if result := Init(NEWRELIC_LICENSE, TEST_APP_NAME); result != 0 {
		t.Errorf("Init test failed. Return: %d.", result)
	}
}

func TestRequestShutdown(t *testing.T) {
    if result := RequestShutdown("Reason: test nr_request_shutdown() "); result != 0 {
		t.Errorf("Shutdown test failed. Return: %d.", result)
	}
}

var shutdownFlag = false
func TestRegisterShutdownCallback(t *testing.T) {
	RegisterShutdownCallback(func() { shutdownFlag = true; fmt.Printf("Shutdown callback called!\n") })

	Init(NEWRELIC_LICENSE, TEST_APP_NAME)
    if result := RequestShutdown("Reason: callback test"); result != 0 {
		t.Errorf("Shutdown test failed. Return: %d.", result)
	}
	if shutdownFlag == false {
		t.Errorf("Shutdown test failed. Callback was not called")
	}
}

func TestDefaultWebTransactionHandler(t *testing.T) {
	DefaultWebTransactionHandler("test default transaction", 0.23)
}


