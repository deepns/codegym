from pathlib import Path
import pprint
import ssl
import socket

# using from C files temporarily until I have the python
# mtls server out.s
CERTDIR = Path("../c/certs")
CERT, KEY, CAFILE = CERTDIR/"client.crt", CERTDIR/"client.key", CERTDIR/"root_ca.crt"
HOST = "localhost"
PORT = 9899

def show_certs(conn: ssl.SSLSocket):
    print(f'Cert from {conn.getpeername()}')
    pprint.pprint(conn.getpeercert())

def create_client_context(certfile, keyfile, cafile) -> ssl.SSLContext:
    ctx = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
    ctx.load_cert_chain(certfile=certfile, keyfile=keyfile)
    ctx.load_verify_locations(cafile=cafile)
    return ctx

def run_client():
    # PROTOCOL_TLS_CLIENT automatically enables the certificate
    # verification and hostname check.
    # context.verify_mode = ssl.CERT_REQUIRED
    # context.check_hostname = False
    context = create_client_context(CERT, KEY, CAFILE)
    # create a TCP socket
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
        # wrap the socket with SSL context
        with context.wrap_socket(sock, server_hostname="localhost") as sslsock:
            # do SSL Connect to connect to the server
            sslsock.connect((HOST, PORT))
            print(f'proto={sslsock.proto}, \
                session={sslsock.session}, \
                negotiated TLS version={sslsock.version()}')
            print(f'Connected from {sslsock.getsockname()} to {sslsock.getpeername()}')
            show_certs(sslsock)

            sslsock.sendall(b'Knock, Knock!')
            data = sslsock.recv(1024).decode('utf-8')
            print(f'Server sent: {data}')

if __name__ == "__main__":
    run_client()