# Daily Dose of Code

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
