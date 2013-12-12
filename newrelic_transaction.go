package newrelic_go_api

/*
#cgo CFLAGS: -fopenmp -I./_include
#cgo LDFLAGS: -lnewrelic-collector-client
#cgo LDFLAGS: -lnewrelic-common
#cgo LDFLAGS: -lnewrelic-transaction

#include <stdlib.h> 
#include <newrelic_transaction.h> 
#include <newrelic_collector_client.h> 
static void setupEmbededCollectorCGOProxy() {
    nr_setup_embedded_collector_client(nr_default_web_transaction_handler);
}

*/
import "C"

import (
	"unsafe"
)

type TTransactionId int64

const (
	NR_ERROR_CODE_OK                      = 0
	NR_ERROR_CODE_INVALID_ID              = 1
	NR_ERROR_CODE_TRANSACTION_NOT_STARTED = 2
	NR_ERROR_CODE_TRANSACTION_IN_PROGRESS = 3
	NR_ERROR_CODE_TRANSACTION_NOT_NAMED   = 4

	NR_DATASTORE_OPERATION_SELECT = "select"
	NR_DATASTORE_OPERATION_INSERT = "insert"
	NR_DATASTORE_OPERATION_UPDATE = "update"
	NR_DATASTORE_OPERATION_DELETE = "delete"
)

func StartWebTransaction() TTransactionId {
	result := C.nr_start_web_transaction()
	return TTransactionId(result)
}

func NameWebTransaction(transactionId TTransactionId, name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	result := C.nr_name_web_transaction(C.long(transactionId), cName)
	return int(result)
}

func EndWebTransaction(transactionId TTransactionId) int {
	result := C.nr_end_web_transaction(C.long(transactionId))
	return int(result)
}

/// Only call once and after nr_init
func SetupEmbeddedCollectorClient() {
	C.setupEmbededCollectorCGOProxy()
}
