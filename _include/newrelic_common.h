/*
 * This is the C interface.
 */

#ifndef NEWRELIC_COMMON_H_
#define NEWRELIC_COMMON_H_

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

static int const NR_LOG_LEVEL_TRACE = 0;
static int const NR_LOG_LEVEL_DEBUG = 1;
static int const NR_LOG_LEVEL_INFO = 2;
static int const NR_LOG_LEVEL_WARNING = 3;
static int const NR_LOG_LEVEL_ERROR = 4;
static int const NR_LOG_LEVEL_FATAL = 5;
static int const NR_LOG_LEVEL_OFF = 6;

static int const NR_ERROR_CODE_INVALID_LOG_FILE_NAME = 1;
static int const NR_ERROR_CODE_INVALID_LOG_FILE_PERMISSIONS = 2;
static int const NR_ERROR_CODE_INVALID_LOG_LEVEL = 10;

int nr_setup_logging(int log_level, const char *log_file_name);

#ifdef __cplusplus
} //! extern "C"
#endif /* __cplusplus */

#endif /* NEWRELIC_COMMON_H_ */
