# What's happening in the garage?

## Things to try out

- [x] Nonblocking TCP client/server
- [-] Running zookeeper cluster locally with docker
- [x] poll
- [x] epoll
- [x] async IO
  - [x] basic signal handling
  - [x] handling SIGIO with sigaction
- [x] Connect to ubuntu VM from vscode using the remote-ssh extension
- [ ] Golang
  - [x] http get
  - [x] http post
  - [x] TLS Server
  - [x] functions
    - [x] variadic
    - [x] closures
    - [x] function types
  - [x] interfaces
    - [x] basics
    - [x] empty interfaces
    - [x] nil interfaces
    - [x] type assertions
  - [x] struct embedding
  - [x] channels
    - [x] basics
    - [x] read only and write only channels
  - [x] select
  - [x] atomic
    - [x] mutex
    - [x] Add/Load/Store
  - [x] environment variables
  - [ ] runes
  - [x] panic/recover
  - [x] strings
    - [x] strings package
    - [x] string formatting
    - [x] string builder
    - [x] string conversions
  - [x] cmdline processing
    - [x] functionality similar to argparse
  - [x] Files
    - [x] temp files
    - [x] directories
  - [x] defer
  - [x] log package
  - [x] init()
  - [x] timers
  - [x] tickers
  - [x] waitgroups
  - [ ] Language and other open source references for go topics
    - [ ] arrays
    - [ ] maps
    - [ ] make vs new
    - [ ] embedding
    - [x] variadic functions
    - [x] interfaces
    - [ ] tempfiles
    - [ ] custom logging
- [ ] syslog, journalctl
- [x] gc pause
- [ ] json path

## Daily log - attempt#2

Been a while I lost in touch with my daily exercise. Restarting the practice.

## Day 25 (grpc interceptor)

- Added a server side unary interceptor
- learnt about how `net.Listen` binds
- came to know about [twirp](https://twitchtv.github.io/twirp/docs/intro.html) - Twirp is similar to gRPC, but without the custom HTTP server and transport implementations: it runs on the standard library's extremely-well-tested-and-high-performance net/http Server. It can run on HTTP 1.1, not just http/2, and supports JSON serialization for easy debugging.
- [ ] can use cURL with gRPC?
- Some links to read
  - [ ] [gRPC use cases](https://www.redhat.com/architect/grpc-use-cases)
  - [ ] [messaging patterns](https://www.redhat.com/architect/architectural-messaging-patterns)
  - [ ] [async message patterns](https://blogs.mulesoft.com/api-integration/patterns/asynchronous-messaging-patterns/)

## Day 24 (grpc metadata, interceptor, realworld examples)

- was wondering how the headers and trailers are used. [this cloudfare article](https://blog.cloudflare.com/road-to-grpc/) had some clear explanation related to that. Usage of trailers in http2 is not so common, grpc uses it to **send status** after the body. e.g. in case of streaming rpcs, **send a checksum** of the content at the end for the client to verify. For streaming data, can't send checksum in the headers since we may not know all the data ahead.
- ![this image](https://blog.cloudflare.com/content/images/2020/10/image3-24.png) neatly summarizes.
- Found these mulesoft articles related to real world use cases. interesting read.
  - [How Gitlab uses grpc](https://www.mulesoft.com/api-university/how-gitlab-puts-grpc-real-world)
  - [How Kubernetes CRI uses grpc](https://www.mulesoft.com/api-university/grpc-real-world-kubernetes-container-runtime-interface)
- read more about interceptors..added a client side example for UnaryInterceptor.
- next?
  - [x] server side unary interceptor
  - [-] client side streaming interceptor

## Day 23 (grpc metadata)

- explored more on metadata, where it is used and why
- see [notes](../go/learning/grpc/features/metadata/readme.md)

## Day 22 (grpc metadata)

- Learning about metadata handling in grpc
- added a sample echo client/service..sending and receiving some key/value pairs in header and trailer
- the header and trailer goes in the respective frames in http/2. how does http/2 frame look like? something to explore next.
- [x] what are the use cases of custom metadata? - one common use case is authentication

## Day 21 (grpc authentication)

- looked into adding authentication to grpc calls
- client and server can be configured to provide authentication using OAuth2.
- client can be configured to use same token for all rpc over the same connection or use token per rpc
- server needs to intercept the call (requires the use of interceptor), validate the token (provided in the metadata) and then return the handler to the original call.
  - This requires learning about metadata handling and interceptor pattern.
- next?
  - [x] metadata - in progress
  - [ ] interceptor

## Day 20 (skip)

## Day 19 (exercise | sliding window maximum)

- picked up [this problem](https://leetcode.com/problems/sliding-window-maximum/), but didn't finish it

## Day 18 (grpc keepalive continued)

- quick recap of basic grpc client and server
- server
  - create a server type with embedded type of the unimplemented server
  - implement the necessary rpc with the server type
  - create a listener on a chosen address and port
  - choose the server options (e.g TLS, KeepAliveParams, KeepAliveEnforcementPolicy etc.)
  - create new server `grpc.NewServer` and register with the service
  - start serving
- client
  - get the server address (host:port)
  - choose the dial options (e.g. tls credentials, keepalive params)
  - dial to the server address with the chosen options. get the connection
  - create a client on the connection returned by dial
  - call the rpc on the client object with the necessary parameters
- keepalive
  - both client and server sends pings...behavior can be controlled by ClientParameters and ServerParameters in `google.golang.org/grpc/keepalive`
  - If no ping, server sends a GOAWAY message to the client to close the connection. I may be missing something here. How the client handles GOAWAY message? I tried to send another message after GOAWAY is received, and the message still went through fine on the same connection. The behavior seem to be library dependent. In general, client is expected to gracefully shutdown the connection upon receiving GOAWAY.
  - http2debug
    - `GODEBUG=http2debug=2` enables verbose logging of the grpc http2 communication. [Doc](https://go.dev/src/net/http/doc.go) says `http2debug=1` also throws some debug logs, but didn't see any in my client testing.
  - ok..didn't know that GOAWAY is a frame type in http2..see [rfc](https://www.rfc-editor.org/rfc/rfc7540#page-43)
  - while looking to find more info about GOAWAY, found [this](https://nuvalence.io/insights/lessons-learned-investigating-goaways-in-grpc-apis/) article. will read through it tomorrow

## Day 17 (grpc keepalive features)

- explored the keepalive features on the client and server side of grpc
- added a client and server application for echo service
- enabled copilot and used it to code the client/server code for keepalive..using it first time...truly game changing experience.

## Day 16 (protobuf serialization)

- added code to serialize protobuf types to pb and json
- explored [keepalive](https://github.com/grpc/grpc-go/blob/master/Documentation/keepalive.md) feature. will add an example shortly.

## Day 15 (exercise|happy number)

- Just picked up [this leetcode problem](https://leetcode.com/problems/happy-number/description/) at random
- recollected some of my forgetten python knowledge.. `//` operator, generators, unittest

## Day 14 (protobuf serialization)

- Explored the possibility of serializing protobuf generated types, in both binary and json format

## Day 13 (grpc)

- Revisit the examples. Updated the service definitions in echo service
- what next?
  - Read the [core concepts](https://grpc.io/docs/what-is-grpc/core-concepts/) again
  - Pick an example feature and explore it

### Day 11, 12 (grpc peer, protobuf definitions)

- explored protobuf [definitions](https://protobuf.dev/programming-guides/proto3/)
- got connection details from context using peer package

### Day 10 (grpc echo server bidirectional streaming)

- added bidirectional streaming rpc to the service
- works, but connection closure is not proper. ~~need to follow up with that tomorrow~~ that error is expected when the stream is closed. Just returned the error.
- TODO - add some comments and clean up the code

### Day 9 (grpc echo server)

- added a client side streaming method to the service
- updated the client and server to use the newly added method

### Day 8 (sick day)

### Day 7 (grpc echo server)

- added a echo server, with proto definition and client and server code
- added a simple echo and server side streaming
- will add client side streaming and bidirectional streaming tomorrow
- read the tutorial page again...finding new information on each read

### Day 6

- went over some patterns from educative

### Day 5 (go contexts)

- More on go contexts
- added a sample program to my learning collection

### Day 4 (go contexts)

- Learning about go contexts

### Day 3 (grpc helloworld, go context)

- Went through the grpc generated code
- Updated the arguments
- started reading about context package...will explore more on that tomorrow.

### Day 2 (grpc helloworld)

- Wrote my own helloworld program on grpc with go
- defined the proto file, explored the marshaling and unmarshalling of proto buffers
- implemented the server and client to test request and response processing
- what next?
  - [x] Read the auto generated files and understand how they are used in general
  - [x] add comments as necessary

### Day 1 (grpc basics)

- running protoc with go requires the extension (protoc-gen-go and protoc-gen-go-grpc). Installed them using `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28` and `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

```console
➜  ~ ls $HOME/go/bin
inputs             protoc-gen-go      protoc-gen-go-grpc tempfiles
➜  ~ export PATH="$PATH:$(go env GOPATH)/bin"
```

- went through the [routeguide example](https://github.com/grpc/grpc-go/tree/master/examples/route_guide), server code and client code. Ran them. what next?
  - [x] create a route guide client on my own to interact with the server
  - [ ] create my own server and client
- server listens at `50051` by default. Unfortunately, that port is already taken by launchd.

```console
➜  ~ sudo lsof -iTCP -sTCP:LISTEN | grep 50051
launchd      1     root   11u  IPv6 0x70998e80006793b3      0t0  TCP *:50051 (LISTEN)
launchd      1     root   50u  IPv4 0x70998e76698d643b      0t0  TCP *:50051 (LISTEN)
```

## Daily log

// Tracking my journey to 100 days and beyond.

- Day 104 #golang
  - Getting started with unit testing
- Day 103 #golang
  - directed channels.
  - recv only channels, send only channels
- Day 102 #golang
  - Ignoring signals
  - registering multiple signals
  - handle signals in go routine
- Day 101 #golang
  - starting to explore signal handling
  - read about os.Signal and os/signal packages, signal registration and handling.
- Day 100 #nocodingday
- Day 99 #golang
  - tried out atomic Add, Load, Store from sync/atomic package
- Day 98 #nocodingday
- Day 97 #nocodingday
- Day 96 #nocodingday
- Day 95 #nocodingday
- Day 94 #golang
  - json encoding of basic types, maps, slices, structs using Marshaller and Encoder
- Day 94 #golang
  - quick recap - part 3 - json decoding, channel send receive, http.Get, go routines, timers, select
- Day 93 #golang
  - quick recap - part 2 - init functions, custom errors, read from stdin, str conversion
  - read using scanner
- Day 92 #golang
  - stable sort
  - quick recap - part 1 - time format, time duration, random generation, nameless structs, slice append, iteration
- Day 91 #golang
  - custom sort functions using sort.Interface interface
- Day 90 #golang
  - sort using functions from sort package. will do custom sorting tomorrow
  - TLS server
- Day 89 #nocodingday
- Day 88 #nocodingday
- Day 87 #nocodingday
- Day 86 #nocodingday
- Day 85 #golang
  - time parsing and formatting
  - converting time to unix epoch seconds and vice versa
- Day 84 #golang
  - explored the usage of WaitGroup from sync package
- Day 83 #golang
  - explored the timers and tickers from time package
- Day 82 #golang
  - multiplex channel read and write with select{..}
  - using time Ticker
- Day 81 #golang
  - channels, bufferred and unbuffered
  - iterate channels
  - added flag processing to ledger example
  - sleep using time.Sleep(). time has the units defined for Second, Milli, Micro
- Day 80 #golang
  - init() functions
- Day 79 #golang
  - flags custom types, duration
- Day 78 #golang
  - flags and args in cmd line
- Day 77 #nocodingday
- Day 76 #nocodingday
- Day 75 #golang
  - continuing sync.Mutex
  - read file line by line using bufio.Scanner
  - added channel to communicate between the workers and main
  - expanded mutex example into a bigger one with multiple go concepts in use.
- Day 74 #golang
  - revisit maps. access missing keys, check key exists or not
  - sync.Mutex. Added a sample ledger app to make use of sync.Mutex
- Day 73 #golang
  - recover, defer
- Day 72 #golang
  - environment variables
  - file metadata
- Day 71 #golang
  - a sample exercise - rot13 subsitution cipher
  - random package and its functions
- Day 70 #golang
  - type assertions
- Day 69 #golang
  - string builder
  - revisited interface, Stringer usage when looking into String() of strings.Builder
  - strconv package. Parsing int,float, bool strings. and reverse conversion using strconv.Format.
- Day 68 #golang
  - nil interfaces
- Day 67 #golang
  - interfaces basics. looked into how fmt.Print functions work.
  - empty interfaces
- Day 66 #golang
  - closures, function variables, anonymous functions
  - struct embedding
- Day 65 #golang
  - variadic functions
  - a self pat 🙌🏻 for reaching 75 days // oops.. spoke too soon.
  - (1) I made a mistake of counting the days.. from 55 I jumped to 66 instead of 56.
  - worked on strings package
- Day 64 #golang
  - tempfiles
  - directories, temp directories
  - defer keyword
- Day 63 #golang
  - constants
- Day 62 #nocodingday
- Day 61 #golang
  - json unmarshaling into defined types and maps
  - http post, url parsing
- Day 60 #golang
  - net/url package
  - net/http Post
- Day 59 #golang
  - errors. creating custom errors.
  - log
    - using log package
    - creating custom loggers
    - logging to files
- Day 58 #golang
  - getting started with error handling
  - creating new errors
- Day 57 #golang
  - json decoding
    - simple values in a json stream, using decoder.Decode()
    - decoding list (using decoder.Token(), More(), Decode())
  - http
    - Get data from URL, decode using a json decoder
- Day 56 #nocodingday
- Day 55 (#golang)
  - #golang - some more structs, file reading, file writing
- Day 54 (#golang)
  - #golang - maps, initialization, iteration
  - #golang - structs, struct methods
- Day 53 (#golang)
  - #golang - pointers, arrays, slices
- Day 52 (#golang)
  - reading from standard input
  - string manipulation, type conversion using strconv, strings package
- Day 51
  - #signal handling - handling return values
  - #signal handling - with sigaction()
  - #signal handling - sigio on non-blocking sockets. This turned out to be a good learning. Tried to read from nonblocking sockets. apparently, sigio tells us only one fd even when multiple file descriptors are ready. I thought a signal would be delivered independently for each fd that is registered with O_ASYNC. I guess we should maintain the list of nonblocking fds to read, when SIGIO is fired, attempt read from all of them.
- Day 50
  - learning #signal handling in C.
- Day 49
  - learning #golang. added some sample code.
  - unclear on how the golang module dependency works. how to import local packages before publishing?
- Day 48 #nocodingday
- Day 47
  - Learning epoll. added a [echo server](../c/epoll/epoll_server_et.c) in edge triggered mode. Been using telnet to talk to the server programs in the last few days.
- Day 46
  - Learning epoll, added an [echo server](../c/epoll/epoll_server.c) in LEVEL Triggered mode (works similar to poll()). will add a separate example for Edge Triggered, with a non-blocking client conn.
  - Didn't add a client this time. Used `telnet <hostname> <IP>` instead.
- Day 45
  - Learnt more on tracing syscalls using **dtruss** in macOS.
- Day 44
  - Added my notes on select in the learning note.
- Day 43
  - Learning about **poll()** system call and its usage. Added a sample [code]([../c/poll/poll_server.c] here)
  - Learning about **non-blocking** sockets. added a simple client/server [demo code](../c/non_blocking/nb_server.c)
  - Didn't know strace is specific to Linux. equivalent tool in macOS is **dtruss** which is a shell script wrapper around **dtrace**
- Day 42 #nocodingday
- Day 41 #nocodingday
- Day 40 #nocodingday
- Day 39 #nocodingday
- Day 38
  - [echo server, multiplexing using select](codegarage/c/select/select_server.c)
- Day 37
  - [echo server client](codegarage/c/echo_server_client/echo_server.c) in C
- Day 36 #nocodingday
- Day 35 #nocodingday
- Day 34
  - [echo server and client](codegarage/python/echo_server_client/echo_server_threaded.py) in python. Also, added a verion of echo server with client processing in separate thread.
- Day 33
  - starting to learn non-blocking sockets
- Day 32
  - #practice problem. [longest-substr-with-ones-after-replacement](practice/longest-subarray-with-ones-after-replacement.py)
- Day 31
  - Deployed a [multiapp cluster](codegarage/kubernetes/multiapp-cluster-with-ingress/deployments/) in minikube. Created multi service with an ingress controller.
- Day 30
  - Deployed a simple webapp on a single node k8s cluster using minikube (see [testapp](codegarage/kubernetes/testapp/README.md))
- Day 29
  - Reading and writing files using io vectors with readv and writev system calls. (sample [here](codegarage/c/vectored_io/vio.c))
  - Modified UDS [server](codegarage/c/domain_sockets/uds_server.c) and [client](codegarage/c/domain_sockets/uds_client.c) to use `read()`/`write()` instead of `send()`/`recv()` calls.
- Day 28
  - Sending messages over unix domain sockets (See [server](codegarage/c/domain_sockets/uds_server.c), [client](codegarage/c/domain_sockets/uds_client.c))
- Day 27
  - Read X509 cert from PEM format file (See [print_cert](codegarage/c/print_cert.c), [cert_file_read](codegarage/c/cert_util.c))
  - Show details about Elliptic Curve key (See [show_cert_info](codegarage/c/cert_util.c)))
- Day 26 - #nocodingday
- Day 25 - #nocodingday
- Day 24 - #nocodingday
- Day 23
  - Continued exploring the SSL cert fields. See [mtls_server](codegarage/c/mtls_server.c)
- Day 22
  - Read SSL cert material sent from the client and print some info about the cert to stdout. See [mtls_server](codegarage/c/mtls_server.c)
- Day 21
  - Using shared memory between two [python scripts](codegarage/python/pydockshm/). Beware of  [resource tracker bug](https://bugs.python.org/issue38119).
  - Sharing shared memory between two python containers
- Day 20
  - [containerized python script](codegarage/python/pydockdemo/Dockerfile)
  - Running simple python scripts in a container without building a docker image.
- Day 19 - #nocodingday
- Day 18
  - Added [mtls client in python](codegarage/python/mtls_client.py)
- Day 17
  - Revisit and fixes to [mtls_server](codegarage/c/mtls_server.c) and [mtls_client](codegarage/c/mtls_client.c).
- Day 16
  - Added [tcp client with mutual TLS](codegarage/c/mtls_client.c)
  - Learnt about difference in TLS handshake behavior between TLS versions 1.2 and 1.3. See the comments in [mtls_client](codegarage/c/mtls_client.c)
- Day 15
  - Added a [tcp server with mutual TLS](codegarage/c/mtls_server.c)
- Day 14
  - Added a [simple tcp server](codegarage/c/tcp_server.c) and [simple tls server](codegarage/c/tls_server.c)
  - Quick refresher of [makefile automatic variables](codegarage/c/Makefile)
- Day 13 (12/10/21)
  - Dump SSL cert info ([python](codegarage/python/ssl_socket_ex.py), [c](codegarage/c/tls_client.c#62))
- Day 12 (12/09/21)
  - [simple tcp client in C](codegarage/c/tcp_client.c)
  - [simple tls tcp client in c](codegarage/c/tls_client.c)
  - [python ssl socket](codegarage/python/ssl_socket_ex.py)
- Day 11 (12/08/21)
  - [continue learning ssl sockets](codegarage/python/ssl_socket_ex.py)
- Day 10 (12/07/21)
  - [Python Socket Programming](codegarage/python/socket_ex.py)
- Day 9
  - [longest-substring-after-k-replacements](practice/length-of-longest-substring.py)
- Day 8
  - [fruits-into-baskets](practice/fruits-into-baskets.py)
  - [longest-substring-with-all-distinct](practice/length-of-longest-substring-with-all-distinct.py)
- Day 7
  - [longest-subarray-with-max-k-distinct](practice/length-of-longest-substring-with-no-more-than-k-distinct-chars.py)
- Day 6
  - [smallest-subarray-with-sum-k](practice/smallest-subarray-with-sum.py)
- Day 5
  - [average-of-contiguous-subarray](practice/find-average-contiguous-subarray.py)
  - [maximum-sum-subarray](practice/maximum-sum-subarray.py)
- Day 4
  - [rearrange-list](practice/rearrange.py)
  - [rearrange-list-max-min](practice/max-min-rearrange.py)
- Day 3
  - [right-rotate](practice/right-rotate.py)
- Day 2 - some practice challenges on python lists
  - [find-product](practice/find-product.py)
  - [find-second-max](practice/find-second-maximum.py)
- Day 1 - Python refresher
  - [list-operations](practice/remove-even-numbers-list.py)
  - [merge-sorted-lists](practice/merge-two-sorted-lists.py)
  - [find-2sum](practice/find-two-numbers-add-upto-k.py)
