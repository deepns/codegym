#! /usr/bin/env python3

# A simple echo server

import argparse
import socket

def handle_connection(conn: socket.socket):
    """ Handler to process client connection """
    data = conn.recv(256).decode('utf-8')
    # getsockname() gets the addr of the local endpoint
    # getpeername() gets the addre of the remote endpoint
    print(f'addr={conn.getpeername()}, data={data}')
    conn.send(data.encode('utf-8'))
    conn.close()

def run_server(host, port, backlog=5):
    """ Start the echo server """
    # using the convenience function create_server which internally
    # does the following steps
    #   create a stream socket
    #     sock.bind((host, port))
    #     sock.listen(5)   
    with socket.create_server((host, port), backlog=backlog) as sock:
        connid = 0
        while True:
            connection, connaddr = sock.accept()
            connid += 1
            print(f'Connected to addr={connaddr}, conn_id={connid}')
            handle_connection(connection)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-a", "--addr", help="host address", type=str, default="localhost")
    parser.add_argument("-p", "--port", help="port number", type=int, default=12345)
    args = parser.parse_args()
    run_server(args.addr, args.port)