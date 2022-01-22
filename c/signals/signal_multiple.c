#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

// Handling multiple signals
// send SIGUSR1 to self when SIGINT is handled
// sending and receiving signals aren't that interesting when
// only one process is involved. will handle signals between
// parent-child processes in future examples.
void signal_handler(int signum) {
    switch(signum) {
        case SIGINT:
        {
            printf("Received SIGINT(%d)\n", signum);
            printf("Sending SIGUSR1\n");
            pid_t pid = getpid();
            kill(pid, SIGUSR1);
            break;
        }
        case SIGUSR1:
            printf("Received SIGUSR1(%d)\n", signum);
            exit(0);
            break;
        default:
            printf("Unknown signal (%d)\n", signum);
            break;
    }
}

int main(void) {
    // register the signal handlers
    signal(SIGINT, signal_handler);
    signal(SIGUSR1, signal_handler);

    for (;;) {
        printf("Sleeping for a sec...\n");
        sleep(1);
    }
    return 0;
}