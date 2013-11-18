package newrelic_go_api

import (
	"testing"
)

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

