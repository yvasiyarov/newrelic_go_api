package newrelic_go_api

/*
#cgo CFLAGS: -fopenmp -I./_include
#cgo LDFLAGS: -lnewrelic-collector-client
#cgo LDFLAGS: -lnewrelic-common
#cgo LDFLAGS: -lnewrelic-transaction

#include <stdlib.h> 
#include <newrelic_collector_client.h> 

extern void statusCallbackCToGoProxy(int status);
static void statusCallbackCGOProxy() {
    nr_register_status_callback(statusCallbackCToGoProxy);
}
*/
import "C"

import (
	"runtime"
	"unsafe"
)

const (
	LANGUAGE = "go"

	NR_STATUS_CODE_SHUTDOWN = 0
	NR_STATUS_CODE_STARTING = 1
	NR_STATUS_CODE_STOPPING = 2
	NR_STATUS_CODE_STARTED  = 3
)

type TShutdownCallback func(int)

func Init(license string, appName string) int {
	cLicense := C.CString(license)
	defer C.free(unsafe.Pointer(cLicense))

	cAppName := C.CString(appName)
	defer C.free(unsafe.Pointer(cAppName))

	cLanguage := C.CString(LANGUAGE)
	defer C.free(unsafe.Pointer(cLanguage))

	cLanguageVersion := C.CString(runtime.Version())
	defer C.free(unsafe.Pointer(cLanguageVersion))

	result := C.nr_init(cLicense, cAppName, cLanguage, cLanguageVersion)
	return int(result)
}

func RequestShutdown(reason string) int {
	cReason := C.CString(reason)
	defer C.free(unsafe.Pointer(cReason))

	result := C.nr_request_shutdown(cReason)
	return int(result)
}

var statusCallback = func(status int) {}

//export statusCallbackCToGoProxy
func statusCallbackCToGoProxy(status C.int) {
	statusCallback(int(status))
}

func RegisterStatusCallback(callback TShutdownCallback) {
	statusCallback = callback
	C.statusCallbackCGOProxy()
}

func DefaultWebTransactionHandler(metricTableJson string) {
	cMetricTableJson := C.CString(metricTableJson)
	defer C.free(unsafe.Pointer(cMetricTableJson))

	C.nr_default_web_transaction_handler(cMetricTableJson)
}
