#! /usr/bin/env python3

import argparse
import socket
import ssl

HOST, PORT = '127.0.0.1', 9090
def run_server():
    context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
    context.load_cert_chain(certfile="certs/server.crt", keyfile="certs/server.key")

    with socket.create_server((HOST, PORT)) as serversock:
        connId = 0
        while True:
            clientsocket, clientaddress = serversock.accept()
            connId += 1
            print(f'Receiving connection from {clientaddress}')
            with context.wrap_socket(clientsocket, server_side=True) as clientsocket_ssl:
                data = clientsocket_ssl.recv(1024).decode('utf-8')
                print(f'Msg from client{connId} {data}')
                clientsocket_ssl.sendall(f'you are conn#{connId}'.encode('utf-8'))

def run_client():

    # use unverified context to connect to server
    # without validating server cert.
    # context = ssl._create_unverified_context()

    context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
    context.load_verify_locations("certs/root_ca.crt")
    context.verify_mode = ssl.CERT_REQUIRED

    # PROTOCOL_TLS_CLIENT automatically enables the hostname verification
    # disabling it explicitly since the certs are generated with self signed
    # root ca
    context.check_hostname = False
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as clientsock:
        with context.wrap_socket(clientsock, server_hostname="127.0.0.1") as clientsock_ssl:
            # upon connect, ssl handshake will be done automatically
            # unless do_handshake_on_connect was set to False when
            # wrapping the socket
            clientsock_ssl.connect((HOST, PORT))

            # send some data to the server
            clientsock_ssl.sendall(b'Hello, server')

            # recv message from server
            data = clientsock_ssl.recv(1024).decode('utf-8')
            print(f'server sent: {data}')

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-m", "--mode", choices=["server", "client"], default="server",
                        help="Run as server or client mode. Default mode:server", required=False)
    args = parser.parse_args()

    if args.mode == "server":
        run_server()
    else:
        run_client()
