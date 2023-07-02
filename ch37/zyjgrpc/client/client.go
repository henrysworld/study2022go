package main

import (
	"context"
	"fmt"
	"github.com/henrysworld/study2022go/ch37/pkg/log"
	pb "github.com/henrysworld/study2022go/ch37/zyjgrpc/proto"
	"google.golang.org/grpc"
	"io"
)

func main() {
	//conn, err := grpc.Dial(":8088", grpc.WithInsecure(), grpc.WithBlock())
	//"retryPolicy": {
	//	"MaxAttempts": 4,
	//		"InitialBackoff": ".01s",
	//		"MaxBackoff": ".01s",
	//		"BackoffMultiplier": 1.0,
	//		"RetryableStatusCodes": [ "UNAVAILABLE" ]
	//}
	conn, err := grpc.Dial(":8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}

	defer conn.Close()
	c := pb.NewCometClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	stream, err := c.Broadcast(context.Background(), &pb.BroadcastReq{
		Message: "xxx",
	})

	if err != nil {
		log.Infof("client failed: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Infof("client.ListFeatures failed: %v", err)

		}
		if err != nil {
			log.Infof("client.ListFeatures Other failed: %v", err)
			conn, stream = retry()
		}

		fmt.Println(msg)
	}

	//var limit int64 = 1
	//var offset int64 = 1
	//r, err := c.ListUser(ctx, &pb.ListUserRequest{
	//	Limit:  &limit,
	//	Offset: &offset,
	//})

	//if err != nil {
	//	log.Fatalf("cound not great: %v", err)
	//}
	//
	//log.Infof("TotalCount: %d", r.TotalCount)
	//
	//for _, user := range r.Users {
	//	d, _ := json.Marshal(user)
	//	log.Info(string(d))
	//}

}

func retry() (*grpc.ClientConn, pb.Comet_BroadcastClient) {
	conn, err := grpc.Dial(":8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}
	c := pb.NewCometClient(conn)
	stream, err := c.Broadcast(context.Background(), &pb.BroadcastReq{
		Message: "xxx",
	})

	if err != nil {
		log.Infof("client failed: %v", err)
	}
	return conn, stream
}
