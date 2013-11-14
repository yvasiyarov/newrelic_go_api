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
    "unsafe"
    "runtime"
)

const LANGUAGE = "go"

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
