#include "logger.h"
#include <stdarg.h>
#include <time.h>
#include <sys/time.h>
#include <stdio.h>

/*
 * TODO
 * - [ ] make timestamp in ISO 8601 format
 *  - [ ] add timezone info?
 * - [ ] convert the variable args into one buffer
 * - [ ] build the logger into separate library
 */

static log_level_t loglevel = LOG_INFO;
static char log_timestamp_buffer[LOG_TIMESTAMP_BUFFER_SIZE];
static char log_msg_buffer[LOG_MSG_BUFFER_SIZE];

void set_log_level(log_level_t level)
{
    loglevel = level;
}

static char * cur_timestamp()
{
    struct timeval tv;
    gettimeofday(&tv, NULL);

    struct tm localtime;
    localtime_r(&tv.tv_sec, &localtime);

    strftime(log_timestamp_buffer, LOG_TIMESTAMP_BUFFER_SIZE, "%Y-%m-%dT%T", &localtime);

    return log_timestamp_buffer;
}

void log_message(log_level_t level, char *format, ...)
{
    if (level <= LOG_ERROR && level >= loglevel) {
        (void)snprintf(log_msg_buffer, LOG_MSG_BUFFER_SIZE, "%s", cur_timestamp());

        fprintf(stderr, "%s\n", log_msg_buffer);
    }
}