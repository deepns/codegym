package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/deepns/codegym/go/learning/grpc/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func UnaryEcho(client pb.EchoServiceClient, message string) {
	request := pb.EchoRequest{
		Message: message,
	}

	// timeout if response is not received within two seconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// connect to the service
	response, err := client.UnaryEcho(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to run SimpleEcho. err=%v", err)
	}

	log.Printf("echo: %v", response.Message)
}

// Run with GODEBUG=http2debug=2 to observe settings frames, ping frames and
// GOAWAY messages due to idleness.
// e.g.
// 2023/04/01 23:22:33 http2: Framer 0x14000144000: wrote SETTINGS len=0
// 2023/04/01 23:22:33 http2: Framer 0x14000144000: read SETTINGS len=6, settings: MAX_FRAME_SIZE=16384
// 2023/04/01 23:22:33 http2: Framer 0x14000144000: read SETTINGS flags=ACK len=0
// 2023/04/01 23:22:33 http2: Framer 0x14000144000: wrote SETTINGS flags=ACK len=0
// 2023/04/01 23:22:33 http2: Framer 0x14000144000: read PING flags=ACK len=8 ping="\x02\x04\x10\x10\t\x0e\a\a"
// 2023/04/01 23:22:38 http2: Framer 0x14000144000: read PING len=8 ping="\x00\x00\x00\x00\x00\x00\x00\x00"
// 2023/04/01 23:22:38 http2: Framer 0x14000144000: wrote PING flags=ACK len=8 ping="\x00\x00\x00\x00\x00\x00\x00\x00"
// 2023/04/01 23:22:43 http2: Framer 0x14000144000: read GOAWAY len=8 LastStreamID=2147483647 ErrCode=NO_ERROR Debug=""
// 2023/04/01 23:22:43 http2: Framer 0x14000144000: read PING len=8 ping="\x01\x06\x01\b\x00\x03\x03\t"
func main() {
	addr := flag.String("addr", "localhost:50055", "The server address in the format of host:port")
	msg := flag.String("msg", "hello", "The message to echo")
	flag.Parse()

	kacs := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             5 * time.Second,
		PermitWithoutStream: true,
	}

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacs),
	}

	conn, err := grpc.Dial(*addr, options...)
	if err != nil {
		log.Fatalf("Failed to dial to %v. err=%v", *addr, err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	UnaryEcho(client, *msg)

	// Wait for 20 seconds to observe GOAWAY due to idleness.
	select {
	case <-time.After(20 * time.Second):
		log.Println("Done")
	}
}
