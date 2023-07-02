package strategytest

import "fmt"

type CSVProcessor struct {
}

func NewCSVProcessor() CSVProcessor {
	return CSVProcessor{}
}

func (c CSVProcessor) Process(oriData any) (newData any) {
	fmt.Printf("csv Process")
	req := oriData.(CSVReq)
	return &CSVResp{
		Ret:  req.Name,
		Data: req.Body,
	}
}

func (c CSVProcessor) Save(data any) error {
	fmt.Printf("csv Save")
	return nil
}

type CSVReq struct {
	Name string
	Body int
}

type CSVResp struct {
	Ret  string
	Data int
}

var _ DataProcessor = (*CSVProcessor)(nil)
