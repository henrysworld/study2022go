package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"testing"
)

const (
	Topic    = "TestZYJ"
	Endpoint = "192.168.0.7:9876"
	//Endpoint  = "0.0.0.0:9876"
	AccessKey = "zyj"
	SecretKey = "zyj123456"
)

func TestCreateTopic(t *testing.T) {
	topic := "test"
	//clusterName := "DefaultCluster"
	nameSrvAddr := []string{Endpoint}
	brokerAddr := "192.168.0.7:10911"

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		fmt.Println(err.Error())
	}

	//create topic
	err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}

	////deletetopic
	//err = testAdmin.DeleteTopic(
	//	context.Background(),
	//	admin.WithTopicDelete(topic),
	//	//admin.WithBrokerAddrDelete(brokerAddr),
	//	//admin.WithNameSrvAddr(nameSrvAddr),
	//)
	if err != nil {
		fmt.Println("Delete topic error:", err.Error())
	}

	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}

func TestProducer(t *testing.T) {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{Endpoint})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	topic := "test"

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
