package main

import (
	"fmt"
)

// 消息发送器接口
type MessageSender interface {
	SendMessage(payload interface{})
}

// 飞书消息发送器
type FeishuMessageSender struct{}

func (f *FeishuMessageSender) SendMessage(payload any) {
	// 发送飞书消息的代码...
	payload = payload.(FeishuPayload)

	fmt.Println("Sending Feishu message:", payload)
}

// 邮件消息发送器
type EmailMessageSender struct{}

func (e *EmailMessageSender) SendMessage(payload any) {
	// 发送邮件的代码...
	fmt.Println("Sending email:", payload)
}

// 短信消息发送器
type SMSMessageSender struct{}

func (s *SMSMessageSender) SendMessage(payload any) {
	// 发送短信的代码...
	fmt.Println("Sending SMS:", payload)
}

// FeishuPayload 结构体包含十个参数
type FeishuPayload struct {
	Param1  string
	Param2  int
	Param3  bool
	Param4  float64
	Param5  string
	Param6  int
	Param7  bool
	Param8  float64
	Param9  string
	Param10 int
}

// EmailPayload 结构体包含七个参数
type EmailPayload struct {
	Param1 string
	Param2 int
	Param3 bool
	Param4 float64
	Param5 string
	Param6 int
	Param7 bool
}

// SMSPayload 结构体包含七个参数
type SMSPayload struct {
	Param1 string
	Param2 int
	Param3 bool
	Param4 float64
	Param5 string
	Param6 int
	Param7 bool
}

// 消息策略接口
type MessageStrategy interface {
	Send(payload interface{})
}

// 策略一：使用飞书发送消息
type StrategyOne struct {
	sender MessageSender
}

func (s *StrategyOne) Send(payload interface{}) {
	s.sender.SendMessage(payload)
}

// 策略二：使用邮件发送消息
type StrategyTwo struct {
	sender MessageSender
}

func (s *StrategyTwo) Send(payload interface{}) {
	s.sender.SendMessage(payload)
}

// 策略三：使用短信发送消息
type StrategyThree struct {
	sender MessageSender
}

func (s *StrategyThree) Send(payload interface{}) {
	s.sender.SendMessage(payload)
}

// 消息策略工厂
type MessageStrategyFactory struct{}

// 创建策略对象
func (f *MessageStrategyFactory) CreateStrategy(strategyType string) MessageStrategy {
	switch strategyType {
	case "strategyOne":
		return &StrategyOne{sender: &FeishuMessageSender{}}
	case "strategyTwo":
		return &StrategyTwo{sender: &EmailMessageSender{}}
	case "strategyThree":
		return &StrategyThree{sender: &SMSMessageSender{}}
	default:
		return nil
	}
}

func main() {
	// 创建消息策略工厂
	factory := &MessageStrategyFactory{}

	// 创建策略对象
	strategyOne := factory.CreateStrategy("strategyOne")
	strategyTwo := factory.CreateStrategy("strategyTwo")
	strategyThree := factory.CreateStrategy("strategyThree")

	// 创建消息Payload
	feishuPayload := FeishuPayload{
		Param1:  "value1",
		Param2:  42,
		Param3:  true,
		Param4:  3.14,
		Param5:  "value5",
		Param6:  10,
		Param7:  false,
		Param8:  2.718,
		Param9:  "value9",
		Param10: 100,
	}

	emailPayload := EmailPayload{
		Param1: "emailValue1",
		Param2: 20,
		Param3: false,
		Param4: 1.234,
		Param5: "emailValue5",
		Param6: 30,
		Param7: true,
	}

	smsPayload := SMSPayload{
		Param1: "smsValue1",
		Param2: 40,
		Param3: true,
		Param4: 4.567,
		Param5: "smsValue5",
		Param6: 50,
		Param7: false,
	}

	// 使用策略对象发送消息
	strategyOne.Send(feishuPayload)
	strategyTwo.Send(emailPayload)
	strategyThree.Send(smsPayload)
}
