#include "logger.h"

int main()
{
    log_message(LOG_INFO, "%s", "Test String");
    return 0;
}