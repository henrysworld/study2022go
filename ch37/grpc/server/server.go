package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"

	pb "github.com/henrysworld/study2022go/ch37/grpc/proto"
	"github.com/henrysworld/study2022go/ch37/pkg/log"
)

func main() {
	startGRPCServer()

	time.Sleep(time.Hour)
}

type server struct {
	pb.UnimplementedMiniBlogServer
}

func (s *server) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	list := make([]*pb.UserInfo, 0)

	createdAt, _ := time.Parse("2006-01-02 15:04:05.000", time.Now().Format("2006-01-02 15:04:05.000"))
	updatedAt, _ := time.Parse("2006-01-02 15:04:05.000", time.Now().Format("2006-01-02 15:04:05.000"))

	log.Infow(time.Now().Format("2006-01-02 15:04:05.000"))
	log.Infow(createdAt.Format("2006-01-02 15:04:05.000"))
	log.Infow(updatedAt.Format("2006-01-02 15:04:05.000"))
	user := &pb.UserInfo{
		Username:  "henry",
		Nickname:  "henry",
		Email:     "henry.zuoyejia.com",
		Phone:     "1878227823",
		PostCount: 100,
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}
	list = append(list, user)
	return &pb.ListUserResponse{
		TotalCount: 0,
		Users:      list,
	}, nil
}

// var _ pb.MiniBlogServer = server{}
var _ pb.MiniBlogServer = (*server)(nil)

//func (s *server) ListUser(context.Context, *pb.ListUserRequest) (*pb.ListUserResponse, error) {
//
//}

func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalw("failed to listen", "err", err)
	}

	grpcsrv := grpc.NewServer()
	pb.RegisterMiniBlogServer(grpcsrv, &server{})

	// 运行 GRPC 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 GRPC 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on grpc address", "addr", ":8080")

	go func() {
		if err := grpcsrv.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()

	return grpcsrv
}
