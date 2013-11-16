package newrelic_collector_client

/*
#cgo CFLAGS: -fopenmp -I./_include
#cgo LDFLAGS: -lnewrelic-collector-client
#cgo LDFLAGS: -lnewrelic-common
#cgo LDFLAGS: -lnewrelic-transaction

#include <stdlib.h> 
#include <newrelic_collector_client.h> 
*/
import "C"

import (
	"runtime"
	"unsafe"
)

const LANGUAGE = "go"

type TShutdownCallback func()

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

var ShutdownCallback = func(){}
func RegisterShutdownCallback(callback TShutdownCallback) {
    C.nr_register_shutdown_callback(C.shutdown_callback(unsafe.Pointer(&ShutdownCallback)))
}

func DefaultWebTransactionHandler(name string, duration float64) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

    C.nr_default_web_transaction_handler(cName, C.double(duration))
}
