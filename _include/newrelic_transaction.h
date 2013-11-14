/*
 * This is the C interface.
 */

#ifndef NEWRELIC_TRANSACTION_H_
#define NEWRELIC_TRANSACTION_H_

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

static int const NR_ERROR_CODE_TRANSACTION_OK = 0;
static int const NR_ERROR_CODE_TRANSACTION_NOT_STARTED = 1001;
static int const NR_ERROR_CODE_TRANSACTION_NOT_NAMED = 1002;
static int const NR_ERROR_CODE_TRANSACTION_INVALID_ID = 1003;


typedef void (*web_transaction_handler)(const char *, double);

/**
 * We assume that each web transaction is executed on a single thread, but you
 * can have multiple web transactions occurring simultaneously.
 * Returns a transaction ID.
 */
long nr_start_web_transaction();

int nr_name_web_transaction(long transaction_id, const char *name);

int nr_end_web_transaction(long transaction_id);

/**
 *
 * Note: This should only be called in the case where you are running the
 * collector client in embedded mode. The web transaction handler should
 * point to a function from newrelic_collector_client.h called
 * nr_default_web_transaction_handler.
 *
 * Only call once and after nr_init
 *
 * @param web_transaction_handler a function pointer to a newrelic funcion
 */
void nr_setup_embedded_collector_client(web_transaction_handler web_transaction_handler);

#ifdef __cplusplus
} //! extern "C"
#endif /* __cplusplus */

#endif /* NEWRELIC_TRANSACTION_H_ */
