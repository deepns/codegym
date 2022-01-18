# Learnings

Noting down some learnings along the way

## I/O multiplexing with select and poll

- Both **select()** and **poll()** are system calls to multiplex I/O among multiple file descriptors (which can be files, sockets, pipes, fifo etc.)
- select
  - declare bit sets (for each of read, write, error operations) for list of sockets to monitor
  - call `select(max_fd, read_fd_set, write_fd_set, error_fd_set, NULL /*timeout*/)`
  - here `max_fd` is the value of the maximum file descriptor. The default value can be obtained from `FD_SETSIZE`.`select` would modify the given fdsets with only the bit of the actual file descriptors available for the given op.
  - Use **FD_SET(fd, fdset)** to set a particular fd in the given fdset.
  - Once select returns, walk all the FDs and operate on the ones whose bit is set in the operation fdset we passed to `select()`.
  - The onus is on the user to keep track of the maximum file descriptor of the process. for e.g. say a server is listening to two sockets simultaneously. Along with the server socket and two client connections, we have three in total but cannot set `nfds` to 3. Need to set max(all known fds) + 1
  - expensive due to the need to lookup all file descriptors each time to check for readiness
- poll

## Tracing system calls of a program in macOS

I didn't know strace didn't exist in macOS. The equivalent of strace in macOS is **dtrace**. macOS has a wrapper shell script **dtruss** around dtrace.

```console
~ % which dtrace
/usr/sbin/dtrace
~ % which dtruss
/usr/bin/dtruss
~ % dtruss --help
/usr/bin/dtruss: illegal option -- -
USAGE: dtruss [-acdefholLs] [-t syscall] { -p PID | -n name | command | -W name }

          -p PID          # examine this PID
          -n name         # examine this process name
          -t syscall      # examine this syscall only
          -W name         # wait for a process matching this name
          -a              # print all details
          -c              # print syscall counts
          -d              # print relative times (us)
          -e              # print elapsed times (us)
          -f              # follow children
          -l              # force printing pid/lwpid
          -o              # print on cpu times
          -s              # print stack backtraces
          -L              # don't print pid/lwpid
          -b bufsize      # dynamic variable buf size
   eg,
       dtruss df -h       # run and examine "df -h"
       dtruss -p 1871     # examine PID 1871
       dtruss -n tar      # examine all processes called "tar"
       dtruss -f test.sh  # run test.sh and follow children
~ %
```

## Set sockets in non-blocking mode

TCP sockets operate in blocking mode by default. They can be turned into non-blocking using **fcntl** or **ioctl** system calls. **ioctl** predates **fcntl**, but can be inconsistent on different implementations and does not support all types of descriptors. **fcntl** is portable, and supports more descriptor types.

```c++

    // Set the socket in non-blocking mode using ioctl
    int ON = 1;
    err = ioctl(sockfd, FIONBIO, (char *)&ON);
    if (err < 0) {
        perror("ioctl() failed");
        exit(EXIT_FAILURE);
    }

    int flags = fcntl(sockfd, F_GETFL, 0);
    err = fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);
    if (err < 0) {
        perror("fcntl() failed");
        exit(EXIT_FAILURE);
    }
```

- When accepting connections on a non-blocking socket, **accept()** would fail with **errno** set to **EWOULDBLOCK** if no connections are waiting to connect
- Likewise, **read()/recv()** would fail with **errno** set to **EAGAIN** if there are no data waiting to be read.
- If **write()/send()** unable to write all the given data, it would fail with **errno** set to **EAGAIN** if data could not be written immediately.

### References

- [SO Thread](https://stackoverflow.com/questions/1150635/unix-nonblocking-i-o-o-nonblock-vs-fionbio)
- [Julia Evans](https://jvns.ca/blog/2017/06/03/async-io-on-linux--select--poll--and-epoll/)
- [devarea](https://devarea.com/linux-io-multiplexing-select-vs-poll-vs-epoll/#.X1hNMNZS9Nc)

## Dumping SSL cert info

- `SSL_get_peer_certificate` get cert info from client SSL object.
- `X509_get_issuer_name` && `X509_NAME_print_ex_fp` combination to dump the cert info to the desired file pointer.

```c++
    X509_NAME *issuer = X509_get_issuer_name(cert);
    X509_NAME *subject = X509_get_subject_name(cert);

    printf("Client Cert issued by:\n");
    X509_NAME_print_ex_fp(stdout, issuer, 2 /*indent*/, XN_FLAG_ONELINE);
    printf("\n");
```

- `X509_NAME_oneline` is deprecated. Use `X509_NAME_print_ex_fp` instead.

## Sharing IPC namespace between containers

1. Start the source container with ipc option as **shareable**
2. Start the other container sharing ipc, with the ipc option as **container:source-container-name**

e.g.

```text
# starting first container
docker run --rm --ipc=shareable --name shmserver --volume $PWD:/home python:latest python /home/shm_server.py

# starting second container
docker run --rm --ipc=container:shmserver --name shmclient --volume $PWD:/home python:latest python /home/shm_client.py
```

## Running a standalone python script in a docker container

A quick way to run simple python scripts in a docker container

docker run command help => `docker run [OPTIONS] IMAGE [COMMAND] [ARG]`

`docker run --rm --volume <path_with_scripts_on_host>:<path_to_script_inside_container> <python_image> python <path_to_script>`

e.g.

```text
pydockdemo $ docker run --name pydockdemo --rm --volume $PWD:/home python:latest python /home/app.py
Hello, World
Running Python version 3.10.1 (main, Dec  8 2021, 03:30:49) [GCC 10.2.1 20210110]
pydockdemo $
```

## OpenSSL Init functions

Prior to 1.1, one has to explicitly initialize the SSL library using during init.

```c
SSL_library_init()
SSL_load_error_strings()
OpenSSL_add_ssl_algorithms()
```

From 1.1 onwards, no explicit initialization is needed. It is taken care by the library internally. If explicit initialization is required, `OPENSSL_init_ssl` can be used. See [this doc](https://www.openssl.org/docs/man1.1.1/man3/OPENSSL_init_ssl.html)

## SSL_connect behavior difference between TLSv1.2 and TLSv1.3

There is a slight difference in behavior between TLSv1.2 and TLSv1.3 in terms of validating the certificate during handshake.In TLSv1.3, SSL_connect() would succeed even if the server rejects client certificate (for e.g. client's cert expired, or invalid, or didn't present at all/). 

Subsequently SSL_write() would also succeed even though server had not accepted the connection fully. The buffer will be written to the socket, but server would not read it. However, SSL_read would fail if the server hadn't accepted the connection.

First read from the client application would then fail since the underlying connection is already closed by the server. In case of TLSv1.2, SSL_connect() would fail if there is a failure in TLS handshake. This [issue](https://github.com/openssl/openssl/issues/8500) has some good comments explaining this version difference.

In OpenSSL 1.1.1, usage of individual TLS method is deprecated using a specific version (for e.g. TLSv1_2_client_method()) will raise the deprecated warning during compilation.

```c
    SSL_CTX *context = SSL_CTX_new(TLSv1_2_client_method());
    if (context == NULL) {
        ERR_print_errors_fp(stderr);
        return NULL;
    }
```

## TCP client/server in C

High level flow in a simple TLS client-server app.

- On the server side
  - Create the TCP socket. Use `socket()`
  - Bind the socket to an addr/port.
    - Set the addr and port in `struct sockaddr_in`. for quick tests on local node, can simply use `INADDR_ANY` for the addr field.
    - Use `bind()` to bind the socket with addr. convert the sockaddr_in to `struct sockaddr` and give the length of the address as well.
  - set the socket in listening mode. Use `listen()`
  - accept the client connections using `accept()` which returns a descriptor to the client socket. `accept()` is blocking by default.
  - work with client socket. reads? `recv()`/`read()`. writes? `send()`/`write()`
  - close the client socket.
- On the client side
  - create socket. Use `socket()`
  - set the addr, port, protocol family in `sockaddr_in`
  - `connect()` socket and addr.
  - read/write to the server.
  - `close()` the socket to server.

## python func args initialization

```python
def func_args_init():
    # the argument values is initialized with an empty list when interpreted.
    # subsequent calls to f() will use the same object assigned to values assigned during initialization
    def f(i, values = []):
        values.append(i)
        print(values)
        return values

    # will print [1], [1, 2], [1, 2, 3]
    f(1)
    f(2)
    f(3)

func_args_init()
```

## pretty python with pprint.pprint

Quick way to pretty print built in data types in python.

```python
>>> # data download from 'https://pypi.org/pypi/sampleproject/json'
>>> print(data)
{'author': 'The Python Packaging Authority', 'author_email': 'pypa-dev@googlegroups.com', 'bugtrack_url': None, 'classifiers': ['Development Status :: 3 - Alpha', 'Intended Audience :: Developers', 'License :: OSI Approved :: MIT License', 'Programming Language :: Python :: 2', 'Programming Language :: Python :: 2.6', 'Programming Language :: Python :: 2.7', 'Programming Language :: Python :: 3', 'Programming Language :: Python :: 3.2', 'Programming Language :: Python :: 3.3', 'Programming Language :: Python :: 3.4', 'Topic :: Software Development :: Build Tools'], 'description': 'A sample Python project\n=======================\n\nThis is the description file for the project.\n\nThe file should use UTF-8 encoding and be written using ReStructured Text. It\nwill be used to generate the project webpage on PyPI, and should be written for\nthat purpose.\n\nTypical contents for this file would include an overview of the project, basic\nusage examples, etc. Generally, including the project changelog in here is not\na good idea, although a simple "What\'s New" section for the most recent version\nmay be appropriate.', 'description_content_type': None, 'docs_url': None, 'download_url': 'UNKNOWN', 'downloads': {'last_day': -1, 'last_month': -1, 'last_week': -1}, 'home_page': 'https://github.com/pypa/sampleproject', 'keywords': 'sample setuptools development', 'license': 'MIT', 'maintainer': None, 'maintainer_email': None, 'name': 'sampleproject', 'package_url': 'https://pypi.org/project/sampleproject/', 'platform': 'UNKNOWN', 'project_url': 'https://pypi.org/project/sampleproject/', 'project_urls': {'Download': 'UNKNOWN', 'Homepage': 'https://github.com/pypa/sampleproject'}, 'release_url': 'https://pypi.org/project/sampleproject/1.2.0/', 'requires_dist': None, 'requires_python': None, 'summary': 'A sample Python project', 'version': '1.2.0'}
>>> pprint.pprint(data, indent=2, depth=2)
{ 'author': 'The Python Packaging Authority',
  'author_email': 'pypa-dev@googlegroups.com',
  'bugtrack_url': None,
  'classifiers': [ 'Development Status :: 3 - Alpha',
                   'Intended Audience :: Developers',
                   'License :: OSI Approved :: MIT License',
                   'Programming Language :: Python :: 2',
                   'Programming Language :: Python :: 2.6',
                   'Programming Language :: Python :: 2.7',
                   'Programming Language :: Python :: 3',
                   'Programming Language :: Python :: 3.2',
                   'Programming Language :: Python :: 3.3',
                   'Programming Language :: Python :: 3.4',
                   'Topic :: Software Development :: Build Tools'],
  'description': 'A sample Python project\n'
  ...
>>> pprint.pprint(data, indent=2, depth=1)
{ 'author': 'The Python Packaging Authority',
  'author_email': 'pypa-dev@googlegroups.com',
  'bugtrack_url': None,
  'classifiers': [...],
  'description': 'A sample Python project\n'
                 '=======================\n'
                 '\n'
                 'This is the description file for the project.\n'
                 '\n'
                 'The file should use UTF-8 encoding and be written using '
                 'ReStructured Text. It\n'
                 'will be used to generate the project webpage on PyPI, and '
                 'should be written for\n'
                 'that purpose.\n'
                 '\n'
                 'Typical contents for this file would include an overview of '
                 'the project, basic\n'
                 'usage examples, etc. Generally, including the project '
                 'changelog in here is not\n'
                 'a good idea, although a simple "What\'s New" section for the '
                 'most recent version\n'
                 'may be appropriate.',
  'description_content_type': None,
  'docs_url': None,
  'download_url': 'UNKNOWN',
  'downloads': {...},
```
