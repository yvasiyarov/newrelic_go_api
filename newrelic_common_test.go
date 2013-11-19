package newrelic_go_api

import (
	"testing"
)

func TestSetupLogging(t *testing.T) {
	if result := SetupLogging(NR_LOG_LEVEL_DEBUG, "/tmp/go_log.txt"); result == 0 {
		t.Errorf("SetupLogging failed. Return %d transaction ID", result)
	}
}

