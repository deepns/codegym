#include <signal.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

/*
 * Trying out sigaction
 * behavior of signal() can vary across different UNIX flavors
 * recommended to use sigaction() for portability and advanced
 * handling of the signals.
 */

/**
 * @brief Handler for signal() call
 * 
 * @param signum 
 */
void signal_handler(int signum) {
    switch (signum) {
    case SIGINT:
        printf("Caught SIGINT. Now quitting\n");
        exit(0);
        break;
    default:
        printf("Unhandled signal %d\n", signum);
        break;
    }
}

/**
 * @brief Handler for sigaction call
 * 
 * @param signum 
 * @param info 
 * @param uap 
 */
void sigaction_handler(int signum, siginfo_t *info, void *uap) {

    // signum - signal for which the handler is invoked

    // info - fields in info are set depending on the delivered signals
    // si_signo, si_errno and si_code are set for all signals
    // si_code - indicates the why the signal was sent
    //         - not seem to be used in macOS
    //         - sending SIGINT from the self, sets the si_code as SI_KERNEL
    //           when running in Ubuntu.
    // for SIGPOLL/SIGIO, si_band will be set (not seeing si_fd in the mac version)

    // uap - ucontext pointer contains the signal context information that 
    // was saved on the user-space stack. Generally not used in the handler

    switch (signum) {
    case SIGINT:
        printf("Caught SIGINT!\n");
        printf("si_signo=%d, si_errno=%d, si_code=0x%x, si_pid=%d\n", 
                info->si_signo, info->si_errno, info->si_code, info->si_pid);
        printf("Quitting now...\n");
        exit(0);
        break;
    default:
        printf("Unhandled signal %d\n", signum);
        break;
    }    
}

int main(void) {

    int err = -1;
    struct sigaction oldaction;

    // vscode intellisense is not so intelligent here
    // not catching the right definitions for sigaction
    // when running in ubuntu

    // sigaction(signum, action, oldaction)
    // if action is not NULL, sigaction sets the handler defined in the action
    // for signum.
    // if oldaction is not NULL, sigaction returns the details of last action
    // set for signum, similar to the return value of signal()
    err = sigaction(SIGINT, NULL, &oldaction);
    if (err == -1) {
        perror("sigaction() failed");
    }

    /*
     * fields of sigaction
     *  sa_mask - signals set in the mask will not be raised in the handler
     *  sa_flags - flags to control the signal delivery and handling
     *  union of sa_handler and sa_sigaction - use sa_sigaction if SA_SIGINFO flag is set
     */
    struct sigaction action;
    action.sa_handler = signal_handler;
    
    // register the signal_handler for SIGUSR1
    err = sigaction(SIGUSR1, &action, NULL);
    if (err == -1) {
        perror("sigaction(SIGUSR1) failed");
    }

    /* 
     * register the SIGINT with sa_sigaction
     * when SA_SIGINFO is set, the signal handler signature
     * requires a function of type void func((int, struct __siginfo *, void *))
     * 
     * siginfo provides even finer information on the delivered signal
     */
    action.sa_flags = SA_SIGINFO;
    action.sa_sigaction = sigaction_handler;
    err = sigaction(SIGINT, &action, NULL);
    if (err == -1) {
        perror("sigaction(SIGINT) failed");
    }

    for (;;) {
        printf("Sleeping for a sec...Ctrl-C to quit\n");
        sleep(1);
    }

    return 0;
}