package newrelic_go_api

/*
#cgo CFLAGS: -fopenmp -I./_include
#cgo LDFLAGS: -lnewrelic-collector-client
#cgo LDFLAGS: -lnewrelic-common
#cgo LDFLAGS: -lnewrelic-transaction

#include <stdlib.h> 
#include <newrelic_common.h> 
*/
import "C"

import (
	"unsafe"
)

const (
	NR_LOG_LEVEL_TRACE   = 0
	NR_LOG_LEVEL_DEBUG   = 1
	NR_LOG_LEVEL_INFO    = 2
	NR_LOG_LEVEL_WARNING = 3
	NR_LOG_LEVEL_ERROR   = 4
	NR_LOG_LEVEL_FATAL   = 5
	NR_LOG_LEVEL_OFF     = 6

	NR_ERROR_CODE_INVALID_LOG_FILE_NAME        = 1
	NR_ERROR_CODE_INVALID_LOG_FILE_PERMISSIONS = 2
	NR_ERROR_CODE_INVALID_LOG_LEVEL            = 10
)

func SetupLogging(logLevel int, logFileName string) int {
	cLogFileName := C.CString(logFileName)
	defer C.free(unsafe.Pointer(cLogFileName))

	result := C.nr_setup_logging(C.int(logLevel), cLogFileName)
	return int(result)
}
