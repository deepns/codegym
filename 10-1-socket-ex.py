#! /usr/bin/env python3

# playing with socket programming in python
# trying a simple client/server comm.

import socket
import argparse

def run_server():
    # socket.socket supports context manager from version 3.2 onwards
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as serversock:
        host, port = 'localhost', 9090
        serversock.bind((host, port))
        serversock.listen()

        connId = 0
        while True:
            clientsocket, clientaddress = serversock.accept()
            connId += 1
            print(f'Receiving connection from {clientaddress}')
            with clientsocket:
                data = clientsocket.recv(1024).decode('utf-8')
                print(f'Msg from client{connId} {data}')
                clientsocket.sendall(f'you are conn#{connId}'.encode('utf-8'))

def run_client():
    clientsock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    host, port = 'localhost', 9090
    clientsock.connect((host, port))

    # send some data to the server
    clientsock.sendall(b'Hello, world')

    # recv message from server
    data = clientsock.recv(1024).decode('utf-8')
    print(f'server sent: {data}')

    clientsock.close()

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-m", "--mode", choices=["server", "client"], default="server",
                        help="Run as server or client mode. Default mode:server", required=False)
    args = parser.parse_args()

    if args.mode == "server":
        run_server()
    else:
        run_client()