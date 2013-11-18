package newrelic_go_api

import (
	"testing"
)

func TestStartWebTransaction(t *testing.T) {
	if result := StartWebTransaction(); result == 0 {
		t.Errorf("StartWebTransaction failed. Return 0 transaction ID")
	}
}
