package main

import (
	"context"
	"encoding/json"
	pb "github.com/henrysworld/study2022go/ch37/grpc/proto"
	"github.com/henrysworld/study2022go/ch37/pkg/log"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8088", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalw("Did not connect", "err", err)
	}

	defer conn.Close()
	c := pb.NewMiniBlogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var limit int64 = 1
	var offset int64 = 1
	r, err := c.ListUser(ctx, &pb.ListUserRequest{
		Limit:  &limit,
		Offset: &offset,
	})

	if err != nil {
		log.Fatalf("cound not great: %v", err)
	}

	log.Infof("TotalCount: %d", r.TotalCount)

	for _, user := range r.Users {
		d, _ := json.Marshal(user)
		log.Info(string(d))
	}

}
