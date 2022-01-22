#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

/*
 * Trying out signal handling
 * signal() is simplified version of sigaction()
 * Register a handler to handle a signal with signal() system call.
 * handler can be one of
 *  SIG_DFL // macro to  (void (*)(int))0
 *  SIG_IGN // macro to (void (*)(int))1
 *  user-defined function
 */

void sigint_handler(int signum) {
    printf("Received SIGINT(%d)\n", signum);
    if (signum == SIGINT) {
        printf("Quitting...\n");
        exit(0);
    }
}

int main(void) {
    // register the signal handlers
    signal(SIGINT, sigint_handler);

    // using default handler
    signal(SIGKILL, SIG_DFL);

    // ignoring SIGUSR1
    signal(SIGUSR1, SIG_IGN);

    for (;;) {
        printf("Sleeping for a sec...\n");
        sleep(1);
    }
    return 0;
}