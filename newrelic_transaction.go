package newrelic_go_api

/*
#cgo CFLAGS: -fopenmp -I./_include
#cgo LDFLAGS: -lnewrelic-collector-client
#cgo LDFLAGS: -lnewrelic-common
#cgo LDFLAGS: -lnewrelic-transaction

#include <stdlib.h> 
#include <newrelic_transaction.h> 
*/
import "C"

const (
	NR_ERROR_CODE_TRANSACTION_OK          = 0
	NR_ERROR_CODE_TRANSACTION_NOT_STARTED = 1001
	NR_ERROR_CODE_TRANSACTION_NOT_NAMED   = 1002
	NR_ERROR_CODE_TRANSACTION_INVALID_ID  = 1003
)

func StartWebTransaction() int {
	result := C.nr_start_web_transaction()
	return int(result)
}
