#ifndef NEWRELIC_COLLECTOR_CLIENT_H_
#define NEWRELIC_COLLECTOR_CLIENT_H_

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

static int const NR_STATUS_CODE_SHUTDOWN = 0;
static int const NR_STATUS_CODE_STARTING = 1;
static int const NR_STATUS_CODE_STOPPING = 2;
static int const NR_STATUS_CODE_STARTED = 3;

typedef void (*nr_status_callback)(int);

int nr_init(const char *license, const char *app_name, const char *language, const char *language_version);

int nr_request_shutdown(const char *reason);

void nr_default_web_transaction_handler(const char *metric_table_json);

void nr_register_status_callback(nr_status_callback callback);

#ifdef __cplusplus
} //! extern "C"
#endif /* __cplusplus */

#endif /* NEWRELIC_COLLECTOR_CLIENT_H_ */
