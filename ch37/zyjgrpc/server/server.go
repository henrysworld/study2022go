package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henrysworld/study2022go/ch37/zyjgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/henrysworld/study2022go/ch37/pkg/log"
	pb "github.com/henrysworld/study2022go/ch37/zyjgrpc/proto"
	"golang.org/x/sync/errgroup"
)

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Banji       Banji     `json:"banji" binding:"required"`
}

type Banji struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func router() http.Handler {
	router := gin.Default()
	router.Use(Cors)
	productHandler := newProductHandler()
	//路由分组、中间件、认证
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			//路由匹配
			productv1.GET(":name", productHandler.Get)
		}
	}

	return router
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type UsageM struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type ChoiceM struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
	Index        int32   `json:"index"`
}

type ResponseGPT struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Usage   UsageM    `json:"usage"`
	Choices []ChoiceM `json:"choices"`
}

func Cors(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	//c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	//c.Header("Content-Type", "application/json")
	//c.AbortWithStatus(200)
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

func (p *productHandler) Get(ctx *gin.Context) {

	//body := strings.NewReader()
	//http.Post("https://api.openai.com/v1/chat/completions")

	//	curl --insecure https://api.openai.com/v1/chat/completions \
	//	-H 'Content-Type: application/json' \
	//	-H 'Authorization: Bearer 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' \
	//	-d '{
	//	"model": "gpt-3.5-turbo",
	//		"messages": [{"role": "user", "content": "实现一个HTML显示可以显示Markdown格式"}]
	//}'
	client := &http.Client{}
	body := make(map[string]interface{})
	msgs := make([]Message, 0, 1)
	msg1 := Message{
		Role:    "user",
		Content: ctx.Param("name"),
	}
	msgs = append(msgs, msg1)
	body["model"] = "gpt-3.5-turbo"
	body["messages"] = msgs

	bytesData, _ := json.Marshal(body)

	fmt.Println(string(bytesData))
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	ret, err := io.ReadAll(resp.Body)

	var responseGPT ResponseGPT
	_ = json.Unmarshal(ret, &responseGPT)

	fmt.Println(string(ret))

	srv.Push(ctx.Param("name"))
	//arrsy := getCs(S)
	//
	//fmt.Println(arrsy)
	//
	//p.RLock()
	//defer p.RUnlock()

	//product, ok := p.products[ctx.Param("name")]
	//if !ok {
	//	ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("can not found product %s", ctx.Param("name"))})
	//	return
	//}

	ctx.JSON(http.StatusOK, responseGPT)

}

var srv *server

func main() {
	//一进行多端口
	var eg errgroup.Group
	insecureServer := &http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// secureServer := &http.Server{
	// 	Addr:         ":8443",
	// 	Handler:      router(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			//log.Fatal(err)
		}

		return err
	})

	ch := zyjgrpc.NewChannel()
	srv = &server{
		Channel: *ch,
	}
	grpcServer := startGRPCServer(srv)

	quit := make(chan os.Signal, 1)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Infow("Shutting down server ...")

	grpcServer.GracefulStop()
}

type server struct {
	pb.UnimplementedCometServer
	zyjgrpc.Channel
}

func (s *server) Stream(streamServer pb.Comet_StreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) PushMsg(msgServer pb.Comet_PushMsgServer) error {
	//for {
	//	//req, err := msgServer.Recv()
	//	//if err != nil {
	//	//	return err
	//	//}
	//
	//	msgServer.SendMsg(&pb.PushMsgReq{
	//		Message: "test" + string(uuid.V4),
	//	})
	//
	//	time.Sleep(time.Second)
	//}

	panic("")
}

func (s *server) Broadcast(req *pb.BroadcastReq, broadcastServer pb.Comet_BroadcastServer) error {
	for {
		//req, err := msgServer.Recv()
		//if err != nil {
		//	return err
		//}

		//msgServer.SendMsg(&pb.PushMsgReq{
		//	Message: "test" + string(uuid.V4),
		//})
		fmt.Println("上线客户端：", req.Message)
		select {
		case <-broadcastServer.Context().Done():
			if broadcastServer.Context().Err() == context.Canceled {
				log.Infof("client %s disconnected", req.Message)
				return context.Canceled
			}
		case proto := <-s.Channel.Signal:
			err := broadcastServer.Send(&pb.BroadcastReply{
				Message: proto.Body,
			})
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.OK:
					// noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					log.Infof("client (%s) terminated connection %s", req.Message, err)
					return err
				default:
					log.Infof("failed to send to client (%s): %v", req.Message, s.Err())
					return err
				}
			}

		}

	}
}

func checkConnected(ctx context.Context, uid string) bool {
	if ctx.Err() == context.Canceled {
		log.Infof("client %s disconnected", uid)
		return false
	}

	return true
}

func (s *server) BroadcastGroup(req *pb.BroadcastGroupReq, groupServer pb.Comet_BroadcastGroupServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *server) Groups(req *pb.GroupsReq, groupsServer pb.Comet_GroupsServer) error {
	//TODO implement me
	panic("implement me")
}

var _ pb.CometServer = (*server)(nil)

// CometServer
func startGRPCServer(cometServer *server) *grpc.Server {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalw("failed to listen", "err", err)
	}

	grpcsrv := grpc.NewServer()
	pb.RegisterCometServer(grpcsrv, cometServer)

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
