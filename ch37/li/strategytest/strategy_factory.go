package strategytest

const (
	A = "A"
	B = "B"
)

type StrategyFactory struct {
	StrategyMap map[string]DataProcessor
}

var mp = make(map[string]DataProcessor)

func NewStrategyFactory() StrategyFactory {
	//var factory DataProcessor
	//factory = NewCSVProcessor()
	mp[A] = NewCSVProcessor()
	mp[B] = NewJsonProcessor()
	return StrategyFactory{StrategyMap: mp}
}

func Newaa() string {
	return "aaa"
}

func GetStrategy(name string) DataProcessor {
	v, ok := mp[name]
	if ok {
		return v
	}

	return nil
}
