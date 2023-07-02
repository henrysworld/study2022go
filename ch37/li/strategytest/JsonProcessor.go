package strategytest

import "fmt"

type JsonProcessor struct {
}

func NewJsonProcessor() JsonProcessor {
	return JsonProcessor{}
}

func (j JsonProcessor) Process(oriData any) (newData any) {
	fmt.Printf("json Process")
	return &JsonResp{
		Ret:  "json abdc",
		Data: 777,
		Type: 999,
	}
}

func (j JsonProcessor) Save(data any) error {
	fmt.Printf("json Save")
	return nil
}

type JsonReq struct {
	Name string
	Body int
	Last int
}

type JsonResp struct {
	Ret  string
	Data int
	Type int
}

var _ DataProcessor = (*JsonProcessor)(nil)
