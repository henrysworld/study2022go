package ch

import (
	"fmt"
)

// 消息发送器接口
type MessageSender interface {
	SendMessage(message string)
}

// 飞书消息发送器
type FeishuMessageSender struct{}

func (f *FeishuMessageSender) SendMessage(message string) {
	// 发送飞书消息的代码...
	fmt.Println("Sending Feishu message:", message)
}

// 邮件消息发送器
type EmailMessageSender struct{}

func (e *EmailMessageSender) SendMessage(message string) {
	// 发送邮件的代码...
	fmt.Println("Sending email:", message)
}

// 短信消息发送器
type SMSMessageSender struct{}

func (s *SMSMessageSender) SendMessage(message string) {
	// 发送短信的代码...
	fmt.Println("Sending SMS:", message)
}

// 消息策略接口
type MessageStrategy interface {
	Send(message string)
}

// 策略一：使用飞书发送消息
type StrategyOne struct {
	sender MessageSender
}

func (s *StrategyOne) Send(message string) {
	s.sender.SendMessage(message)
}

// 策略二：使用邮件发送消息
type StrategyTwo struct {
	sender MessageSender
}

func (s *StrategyTwo) Send(message string) {
	s.sender.SendMessage(message)
}

// 策略三：使用短信发送消息
type StrategyThree struct {
	sender MessageSender
}

func (s *StrategyThree) Send(message string) {
	s.sender.SendMessage(message)
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

	// 使用策略对象发送消息
	strategyOne.Send("Hello from Strategy One")
	strategyTwo.Send("Hello from Strategy Two")
	strategyThree.Send("Hello from Strategy Three")
}
