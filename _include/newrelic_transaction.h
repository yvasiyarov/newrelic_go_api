#ifndef NEWRELIC_TRANSACTION_H_
#define NEWRELIC_TRANSACTION_H_

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

static int const NR_ERROR_CODE_OK = 0;
static int const NR_ERROR_CODE_INVALID_ID = 1;
static int const NR_ERROR_CODE_TRANSACTION_NOT_STARTED = 2;
static int const NR_ERROR_CODE_TRANSACTION_IN_PROGRESS = 3;
static int const NR_ERROR_CODE_TRANSACTION_NOT_NAMED = 4;

static const char * const NR_DATASTORE_OPERATION_SELECT = "select";
static const char * const NR_DATASTORE_OPERATION_INSERT = "insert";
static const char * const NR_DATASTORE_OPERATION_UPDATE = "update";
static const char * const NR_DATASTORE_OPERATION_DELETE = "delete";

typedef void (*web_transaction_handler)(const char *);

long nr_start_web_transaction();
int nr_name_web_transaction(long id, const char *name);
int nr_end_web_transaction(long id);
long nr_start_datastore_statement(const char *table, const char *operation);
int nr_end_datastore_statement(long id);

/**
 *
 * Note: This should only be called in the case where you are running the
 * collector client in embedded mode. The web transaction handler should
 * point to a function from newrelic_collector_client.h called
 * nr_default_web_transaction_handler.
 *
 * Only call once and after nr_init.
 *
 * @param web_transaction_handler a function pointer to a newrelic transaction callback
 */
void nr_setup_embedded_collector_client(web_transaction_handler nr_default_web_transaction_handler);

#ifdef __cplusplus
} //! extern "C"
#endif /* __cplusplus */

#endif /* NEWRELIC_TRANSACTION_H_ */
