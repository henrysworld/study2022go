package strategytest

//type IData interface {
//	ProcessAndGetData(data any) any
//}

type DataProcessor interface {
	Process(oriData any) (newData any)
	Save(data any) error
}
