package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}

// func main() {
// 	var wg sync.WaitGroup
// 	var num int = 5

// 	wg.Add(num)

// 	for i := 0; i < num; i++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			startClient(i)
// 		}(i + 1)
// 	}
// 	wg.Wait()
// }

// func startClient(i int) {
// 	quit := make(chan struct{})
// 	done := make(chan struct{})
// 	conn, err := net.Dial("tcp", ":8888")
// 	if err != nil {
// 		fmt.Println("dial error:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	fmt.Printf("[client %d]: dial ok\n", i)

// 	// 生成payload
// 	rng, err := codename.DefaultRNG()
// 	if err != nil {
// 		panic(err)
// 	}

// 	frameCodec := frame.NewMyFrameCodec()
// 	var counter int

// 	go func() {
// 		// handle ack
// 		for {
// 			select {
// 			case <-quit:
// 				done <- struct{}{}
// 				return
// 			default:
// 			}

// 			conn.SetReadDeadline(time.Now().Add(time.Second * 5))
// 			ackFramePayLoad, err := frameCodec.Decode(conn)
// 			if err != nil {
// 				if e, ok := err.(net.Error); ok {
// 					if e.Timeout() {
// 						continue
// 					}
// 				}
// 				panic(err)
// 			}

// 			p, err := packet.Decode(ackFramePayLoad)
// 			submitAck, ok := p.(*packet.SubmitAck)
// 			if !ok {
// 				panic("not submitack")
// 			}
// 			fmt.Printf("[client %d]: the result of submit ack[%s] is %d\n", i, submitAck.ID, submitAck.Result)
// 		}
// 	}()

// 	for {
// 		// send submit
// 		counter++
// 		id := fmt.Sprintf("%08d", counter) // 8 byte string
// 		payload := codename.Generate(rng, 4)

// 		s := &packet.Submit{
// 			ID: id,
// 			// Payload: []byte(payload),
// 			Payload: []byte("hello11111" + payload),
// 		}

// 		framePayload, err := packet.Encode(s)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("[client %d]: send submit id = %s, payload=%s, frame length = %d\n",
// 			i, s.ID, s.Payload, len(framePayload)+4)

// 		err = frameCodec.Encode(conn, framePayload)
// 		if err != nil {
// 			panic(err)
// 		}

// 		time.Sleep(1 * time.Second)
// 		if counter >= 10 {
// 			quit <- struct{}{}
// 			<-done
// 			fmt.Printf("[client %d]: exit ok\n", i)
// 			return
// 		}
// 	}
// }
