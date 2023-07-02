package zyjgrpc

type Channel struct {
	//Room     *Room
	//CliProto Ring
	Signal chan Proto
	//Writer bufio.Writer
	//Reader bufio.Reader
	//Next   *Channel
	//Prev   *Channel
	//
	//Mid      int64
	//Key      string
	//IP       string
	//watchOps map[int32]struct{}
	//mutex sync.RWMutex
}

func NewChannel() *Channel {
	return &Channel{
		Signal: make(chan Proto, 100),
	}
}

func (c *Channel) Push(s string) {

	c.Signal <- Proto{
		Body: s,
	}
}

type Proto struct {
	Body string `json:"body,omitempty"`
}
