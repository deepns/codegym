#! /usr/bin/env python3

# A simple tcp client to talk to a echo server

import argparse
import socket

def run_client(host, port):
    """ Connect to the server and exchange the messages """
    with socket.create_connection((host, port)) as sock:
        print(f'Connected from {sock.getsockname()} to {sock.getpeername()},'
            f'timeout={sock.gettimeout()},'
            f'defaulttimeout={socket.getdefaulttimeout()}')
        sock.send(b'Hello, there')
        msgfromserver = sock.recv(1024).decode('utf-8')
        print(f'msg from server: {msgfromserver}')

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-a", "--addr", help="host address", type=str, default="localhost")
    parser.add_argument("-p", "--port", type=int, help="server port number: defaults to 12345", default=12345)
    args = parser.parse_args()
    run_client(args.addr, args.port)