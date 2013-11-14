/*
 * This is the C interface.
 */

#ifndef NEWRELIC_COLLECTOR_PROXY_H_
#define NEWRELIC_COLLECTOR_PROXY_H_

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

typedef void (*shutdown_callback)();

int nr_init(const char *license, const char *app_name, const char *language, const char *language_version);

int nr_request_shutdown(const char *reason);

void nr_default_web_transaction_handler(const char *name, double duration);

void nr_register_shutdown_callback(shutdown_callback callback);

#ifdef __cplusplus
} //! extern "C"
#endif /* __cplusplus */

#endif /* NEWRELIC_COLLECTOR_PROXY_H_ */
