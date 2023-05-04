#! /usr/bin/env bash

# TODO
# - [ ] take target-directory as input and generate all the files in the target dir
# - [ ] use getopt for argument processing

DAYS_VALID=${DAYS_VALID:-356}

# Generate the private key and self-signed cert for the root CA
generate_root_ca()
{
    ROOT_CA=$1
    # Create a self signed certificate for the root CA
    # what are these options mean?
    # -x509 - output a self signed certificate instead of a certificate request. Commonly used to generate a test certificate or self-signed cert for the root CA.
    # -nodes - stands for No DES (Data Encryption Standard), not nodes as in servers/hosts. Used to not encrypt the private key
    # -sha256 - use sha256 hashing to sign the request // optional
    # -newkey arg - create a new private key along with the request/certificate. "arg" specifies the type and length of the key to generate. The new key is written to standard output.
    # -key <key file> - used if an existing key file needs to be used
    # -days <n> - certificate is valid n number of days.
    # -out <output_file> - write the request/certificate to the output file instead of standard out.
    # -subj <arg> - specificy the subject name of the request. C-Country, ST-state, O-Organization, CN=Common Name.
    #             - openssl will prompt the user to enter the subject fields. Passing them in -subj makes it convenient to script
    openssl req -x509 -nodes -sha256 -newkey rsa:2048 \
        -subj "/C=US/ST=DE/O=ExampleRootCA, Inc./CN=examplerootca.org" \
        -days ${DAYS_VALID} -out $ROOT_CA.crt -keyout $ROOT_CA.key
}

generate_cert_signed_by_root_ca()
{
    ROOT_CA=$1
    NODE=$2

    # Generate the CSR for the server

    # The common name of the node/server is different from the common name of the root CA
    # The common name is typically the address to which we connect to the server
    # Set CN=<IP address|hostname> of the server for which the certificate is being generated.
    openssl req -new -nodes -newkey rsa:2048 \
        -subj "/C=US/ST=DE/O=Example-$2, Inc./CN=localhost" \
        -addext "subjectAltName = DNS:localhost, IP:127.0.0.1" \
        -out $NODE.csr -keyout $NODE.key

    # Generate the certificate signed by the root CA
    # -req input is a CSR
    # -sha256 - use sha256 message digest to sign the request. Default is sha1.
    # -days <n> - certify the certificate for n number of days
    # -in <file> - read the CSR from the input file
    # -CA <file> - path to CA root certificate
    # -CAkey <file> - path to private key of CA
    # -CAcreateserial - create CA serial number file if it doesn't exist already. Don't know where it is used in the certificate
    # -CAserial <file> - use the serial number specified in this file
    # -set_serial <number> - serial number specified in decimal or hexadecimal format. -CAserial and CAcreateserial are ignored when this option is specified
    openssl x509 -req -sha256 -days ${DAYS_VALID} \
        -in $2.csr -CA $ROOT_CA.crt \
        -extfile <(printf "subjectAltName=DNS:localhost,IP:127.0.0.1") \
        -CAkey $ROOT_CA.key -CAcreateserial -out $NODE.crt
}

# Validate the cert is signed by the given private key
# e.g. validate_cert_key server.crt server.key
validate_cert_key()
{
    # Validate the server certificate was signed using the server key.
    # the modulus parameter of the server certificate and the server key
    # should match if the certificate and private key are consistent

    # the output of "openssl rsa -noout -in server.key -modulus"
    # resemble like below
    # Modulus=CFA6C9C89E204EB6DAA396BC1F1A49F07CC96A97E4B09A5C778124D7788294AF782B23EAD9C8966C5D803BEB242C300E52AE99C86EBE0D73B85AA9156DB675296A7828FA7B93AD501CE6884D53C939F0CE046CE41233DF6EBF17211B25248A3B562E5865FAFC3C2F4870E9D4ABAD2FAE41AAE9BEFB7403B414A4C5665204361497F91B606951EFE66E5152221A99C8BA9064EAB073F88BE9454104B6418D6450040176AC277D89F95D947FE58ED305A9F9E62569C034989A817669C77515909F42512C42614273CC8033BC178C4F5464E14EDB395F54257AFEF279539DD83B66E4618DBBB5401DA2540110D9F10B52C4C5A5C467AC8CC3AA7DF011D8CA4042E3

    # modulus length is same as the key length
    # See https://en.wikipedia.org/wiki/RSA_(cryptosystem)#Key_generation
    # for modulus calcuation

    # taking the md5 hash of the modulus makes it easier to compare
    # x509 - Print the value of the modulus of the public key contained in the certificate.
    # rsa - Print the value of the modulus of the key.

    MD5CMD=$(which md5sum 2> /dev/null)
	if [[ $? != 0 ]]; then
		# md5sum equivalent command in freebsd (and macOS) is md5
		# Try using md5 if md5sum is not available
    	MD5CMD=$(which md5)
	fi

    cert_md5=$(openssl x509 -noout -in $1 -modulus | $MD5CMD)
    key_md5=$(openssl rsa -noout -in $2 -modulus | $MD5CMD)

    if [[ $cert_md5 == $key_md5 ]]; then
        echo "Cert and key match"
    else
        echo "Cert and key do not match"
    fi
}

ROOT_CA=$1
generate_root_ca $ROOT_CA

# shift the args to the left.
shift

# for the remaining args, generate a cert signed by the root CA
for node in $@;do
    generate_cert_signed_by_root_ca $ROOT_CA $node
    validate_cert_key $node.crt $node.key
done
