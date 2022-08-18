package main

import (
	"context"
	"log"
	"net"

	pb "github.com/henrysworld/study2022go/ch37/cmd/helloworld"
	"google.golang.org/grpc"
)

var (
	port = ":50052"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed t o serve: v%", err)
	}

}

// func main() {
// 	l, err := net.Listen("tcp", ":8888")
// 	if err != nil {
// 		fmt.Println("listen error:", err)
// 		return
// 	}

// 	fmt.Println("server start ok(on *.8888)")

// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			fmt.Println("accept error:", err)
// 			break
// 		}

// 		go handleConn(c)
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	frameCodec := frame.NewMyFrameCodec()

// 	for {
// 		framePayload, err := frameCodec.Decode(c)
// 		if err != nil {
// 			fmt.Println("handleConn: frame decode error:", err)
// 			return
// 		}

// 		ackFramePayload, err := handlePacket(framePayload)
// 		if err != nil {
// 			fmt.Println("handleConn: handle packet error", err)
// 		}

// 		err = frameCodec.Encode(c, ackFramePayload)
// 		if err != nil {
// 			fmt.Println("handleConn: frame encode error:", err)
// 			return
// 		}
// 	}
// }

// func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
// 	var p packet.Packet
// 	p, err = packet.Decode(framePayload)
// 	if err != nil {
// 		fmt.Println("handleConn: packet decode error:", err)
// 		return
// 	}

// 	switch p.(type) {
// 	case *packet.Submit:
// 		submit := p.(*packet.Submit)
// 		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
// 		submitAck := &packet.SubmitAck{
// 			ID:     submit.ID,
// 			Result: 0,
// 		}

// 		ackFramePayload, err = packet.Encode(submitAck)
// 		if err != nil {
// 			fmt.Println("handleConn: packet encode error:", err)
// 			return nil, err
// 		}

// 		return ackFramePayload, nil

// 	default:
// 		return nil, fmt.Errorf("unknow packet type")
// 	}
// }
