package newrelic_collector_client

import (
    "testing"
)

func TestInit(t *testing.T) {

    if result := Init("7bceac019c7dcafae1ef95be3e3a3ff8866de246", "TestApp"); result != 0 {
        t.Errorf("Init tes failed. Return: %d.", result)
    }
}
