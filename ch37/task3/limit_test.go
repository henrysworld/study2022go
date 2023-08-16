package task3

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sort"
	"testing"
	"time"
)

var ctx = context.Background()

func isAllowed(client *redis.Client, userId string, actionKey string, period time.Duration, maxCount int) (bool, error) {
	key := fmt.Sprintf("hist:%s:%s", userId, actionKey)

	now := time.Now().UnixNano()
	// 设定时间窗口
	windowStart := now - int64(period)

	// 使用事务进行操作，确保原子性
	pipe := client.TxPipeline()

	// 记录当前时间戳
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})

	// 移除时间窗口之前的记录
	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

	// 获取窗口内的记录数量
	countCmd := pipe.ZCard(ctx, key)

	// 设置key的过期时间以自动清理，防止无限增长
	pipe.Expire(ctx, key, period+time.Second)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	// 如果在时间窗口内的请求超过了最大次数，则拒绝
	return countCmd.Val() <= int64(maxCount), nil
}

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

var client *redis.Client

func TestLimit(t *testing.T) {
	client = NewClient()
	defer client.Close()
	time.Sleep(10 * time.Second)
	userId := "12345678"
	action := "view_page"
	maxCount := 5
	period := 1 * time.Minute

	for i := 0; i < 1000; i++ {
		ix := i
		go func() {
			allowed, err := isAllowed(client, userId, action, period, maxCount)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if allowed {
				fmt.Println(ix, "===Request allowed!")
			} else {
				fmt.Println(ix, "===Request denied!")
			}
		}()

		//time.Sleep(10 * time.Second)
	}

	time.Sleep(100 * time.Second)
}

func TestName111(t *testing.T) {
	cw := &CtrlWork{
		Times: []string{"2023-06-27 00:43:22", "2023-06-28 00:43:22", "2023-06-26 00:43:22", "2023-06-18 00:43:22", "2023-06-29 00:43:22"},
		Temps: []float32{56.89, 87.23, -1, -1, 87.12},
	}

	//sort.Slice(cw.Times, func(i, j int) bool {
	//	return cw.Times[i] > cw.Times[j]
	//})
	//
	//fmt.Println(cw.Times)
	//fmt.Println(cw.Temps)
	sort.Sort(&ctrlWorkSorter{cw})

	fmt.Println("Sorted Times:", cw.Times)
	fmt.Println("Corresponding Temps:", cw.Temps)
}

type CtrlWork struct {
	Times []string
	Temps []float32
}

type ctrlWorkSorter struct {
	cw *CtrlWork
}

func (c *ctrlWorkSorter) Len() int {
	return len(c.cw.Times)
}

func (c *ctrlWorkSorter) Less(i, j int) bool {
	return c.cw.Times[i] > c.cw.Times[j] // 降序排序
}

func (c *ctrlWorkSorter) Swap(i, j int) {
	c.cw.Times[i], c.cw.Times[j] = c.cw.Times[j], c.cw.Times[i]
	c.cw.Temps[i], c.cw.Temps[j] = c.cw.Temps[j], c.cw.Temps[i]
}

//func (c *CtrlWork) Len() int {
//	return len(c.Times)
//}
//
//func (c *CtrlWork) Less(i, j int) bool {
//	return c.Times[i] > c.Times[j]
//}
//
//func (c *CtrlWork) Swap(i, j int) {
//	c.Times[i], c.Times[j] = c.Times[j], c.Times[i]
//	c.Temps[i], c.Temps[j] = c.Temps[j], c.Temps[i]
//}
