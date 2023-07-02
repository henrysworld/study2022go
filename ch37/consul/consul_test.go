package consul

import (
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"testing"
)

func TestName(t *testing.T) {
	cfg := capi.DefaultConfig()
	cfg.Address = "localhost:8500"

	client, err := capi.NewClient(cfg)

	if err != nil {
		panic(err)
	}

	srv := NewService(client)
	config := NewConfig(client)

	err = srv.Register()
	if err != nil {
		panic(err)
	}
	data, err := srv.AllServiceList()
	if err != nil {
		panic(err)
	}

	config.Get("REDIS_MAXCLIENTS")

	for key, _ := range data {
		fmt.Println(key)
	}
}

type Service struct {
	client *capi.Client
	Name   string `json:"name"`

	IP      string   `json:"ip"`
	Port    int      `json:"port"`
	Address string   `json:"address"`
	Tags    []string `json:"tags"`
}

type Config struct {
	client *capi.Client
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func NewService(client *capi.Client) *Service {

	return &Service{
		client: client,
	}
}

func NewConfig(client *capi.Client) *Config {
	return &Config{
		client: client,
	}
}

func (c *Config) Create(key, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *Config) Get(key string) (string, error) {

	//2.配置中心
	kv := c.client.KV()
	pair, mate, err := kv.Get("REDIS_MAXCLIENTS", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("KV: %s %s\n", pair.Key, pair.Value)
	fmt.Printf("QueryMeta: %v\n", mate)

	return string(pair.Value), nil
}

func (c *Config) Update(key, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *Config) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Register() error {

	//生成对应的检查对象
	//check := &capi.AgentServiceCheck{
	//	HTTP:                           "http://localhost:8021/health",
	//	Timeout:                        "5s",
	//	Interval:                       "5s",
	//	DeregisterCriticalServiceAfter: "10s",
	//}
	//生成注册对象
	registration := new(capi.AgentServiceRegistration)
	registration.Name = "sse-conn-comet"
	registration.ID = "1"
	registration.Port = 9443
	registration.Tags = []string{"sse-conn-comet"}
	registration.Address = "192.168.0.8"
	//registration.Check = check

	err := s.client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *Service) AllServiceList() ([]*Service, error) {
	data, err := s.client.Agent().Services()
	if err != nil {
		panic(err)
	}

	var keys []string
	for key, _ := range data {
		keys = append(keys, key)
		fmt.Println(key)
	}

	ret := make([]*Service, len(keys))

	fmt.Printf("data: %+v\n", data)
	for i := 0; i < len(data); i++ {
		ret[i] = &Service{
			Name: data[keys[i]].Service,
		}
	}

	return ret, nil
}

func (s *Service) Find(name string) ([]*Service, error) {
	data, err := s.client.Agent().ServicesWithFilter(name)
	if err != nil {
		panic(err)
	}

	keys := make([]string, len(data))
	for key, _ := range data {
		keys = append(keys, key)
		fmt.Println(key)
	}

	ret := make([]*Service, len(keys))

	fmt.Printf("data: %+v\n", data)
	for i := 0; i < len(data); i++ {
		ret[i] = &Service{
			Name: data[keys[i]].Service,
		}
	}

	return ret, nil
}

func (s *Service) DeRegister(id string) error {
	//TODO implement me
	panic("implement me")
}

type IService interface {
	Register() error
	AllServiceList() ([]*Service, error)
	Find(name string) ([]*Service, error)
	DeRegister(id string) error
}

type IConfig interface {
	Create(key, value string) error
	Get(key string) (string, error)
	Update(key, value string) error
	Delete(key string) error
}

var _ IService = (*Service)(nil)

var _ IConfig = (*Config)(nil)
