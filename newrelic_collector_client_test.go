package newrelic_collector_client

import (
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
	if result := Init(NEWRELIC_LICENSE, TEST_APP_NAME); result != 0 {
		t.Errorf("Init test failed. Return: %d.", result)
	}
	if result := RequestShutdown("shutdown reason"); result != 0 {
		t.Errorf("Shutdown test failed. Return: %d.", result)
	}
}
