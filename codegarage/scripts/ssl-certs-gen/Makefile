ROOT_CA=root_ca
NODES=server client

certs:
	./generate_ssl_certs.sh ${ROOT_CA} ${NODES}
clean:
	rm -vf *.key *.srl *.crt *.csr
