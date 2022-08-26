typedef enum {
    LOG_DEBUG,
    LOG_INFO,
    LOG_WARNING,
    LOG_ERROR
} log_level_t;

#define LOG_TIMESTAMP_BUFFER_SIZE 512
#define LOG_MSG_BUFFER_SIZE 4096

void set_log_level(log_level_t level);

void log_message(log_level_t level, char *format, ...);