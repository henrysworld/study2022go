package strategytest

import (
	"fmt"
	"testing"
)

func TestProc(t *testing.T) {
	//fa := NewStrategyFactory()
	//fa := Newaa()
	//fmt.Printf("%v", fa)
	a, ok := NewStrategyFactory().StrategyMap["A"]
	if ok {
		data := a.Process(CSVReq{
			Name: "a",
			Body: 666,
		})

		fmt.Printf("ch:%s", data)
	}
}

func TestNode(t *testing.T) {
	fmt.Printf("ssssss")
}
