package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 这应该是你的 Redis 服务器地址
		Password: "",               // 如果没有密码，就留空
		DB:       0,                // 使用默认 DB
	})

	ctx := context.Background()

	// 删除 Redis 中的 counter 键，以防止之前的运行影响这次的结果
	rdb.Del(ctx, "counter")

	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 使用 Redis 的 INCR 命令增加 counter 的值
			val, err := rdb.Incr(ctx, "counter").Result()
			if err != nil {
				log.Fatal(err)
			}

			// 如果 counter 的值小于或等于 10，向 sync.Map 添加键值对
			if val <= 10 {
				m.Store(i, i)
				fmt.Printf("Inserted item %d into the map.\n", i)
			}
		}(i)
	}

	wg.Wait()
	js, _ := SyncMapToJSON(&m)
	fmt.Printf("map: %s\n", js)
}

func SyncMapToJSON(m *sync.Map) (string, error) {
	// 将 sync.Map 转换为普通的 map
	var normalMap = make(map[int]interface{})
	m.Range(func(key, value interface{}) bool {
		normalMap[key.(int)] = value
		return true
	})

	// 将 map 转换为 JSON
	jsonData, err := json.Marshal(normalMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
