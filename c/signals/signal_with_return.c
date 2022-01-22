#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

/*
 * Trying out signal handling
 * signal() is simplified version of sigaction()
 * Register a handler to handle a signal with signal(signum, handler) system call.
 * handler can be one of
 *  SIG_DFL // macro to  (void (*)(int))0
 *  SIG_IGN // macro to (void (*)(int))1
 *  user-defined function
 * signal(signum, action) returns the last action set for the given signum
 */

/**
 * @brief handler to handle SIGNINT
 * 
 * @param signum 
 */
void sigint_handler(int signum) {
    printf("Received SIGINT(%d)\n", signum);
    if (signum == SIGINT) {
        printf("Quitting...\n");
        exit(0);
    }
}

int main(void) {
    void *old_action;

    // register the signal handlers
    // signal(signum, action) returns the last action set for the given signum
    // here, signal() would return SIG_DFL (which maps to 0x0)
    old_action = signal(SIGINT, sigint_handler);
    printf("old_action=%p\n", old_action);

    // registering SIGINT again, but resetting to default.
    // signal() would return the address of sigint_handler
    old_action = signal(SIGINT, SIG_DFL);
    printf("old_action=%p\n", old_action);

    // ignoring SIGUSR1
    // here, signal() would return SIG_DFL (which maps to 0x0)
    // since no handler is registered for SIGUSR1 before
    old_action = signal(SIGUSR1, SIG_IGN);
    printf("old_action=%p\n", old_action);

    // can't ignore SIGKILL. signal() will error out here and return -1.
    old_action = signal(SIGKILL, SIG_IGN);
    printf("old_action=%p\n", old_action);
    if (old_action == SIG_ERR) {
        perror("signal(SIGKILL) failed");
    }

    for (;;) {
        printf("Sleeping for a sec...\n");
        sleep(1);
    }
    return 0;
}