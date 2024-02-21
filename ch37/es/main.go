package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7" // 确保使用的是你安装的elastic版本对应的路径
	"log"
	"reflect"
)

func main() {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), // 使用你的Elasticsearch地址
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 使用NewTermsQuery构建查询
	vins := []interface{}{"HLX33B125P1184429", "另一个VIN", "再一个VIN"} // 使用interface{}切片以匹配NewTermsQuery的参数要求
	termsQuery := elastic.NewTermsQuery("vin", vins...)

	// 执行查询
	searchResult, err := client.Search().
		Index("你的索引名称"). // 替换为你的索引名
		Query(termsQuery).
		Sort("report_time", true). // 根据需要进行排序
		From(0).Size(10).          // 分页参数
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	// 打印查询结果
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found a total of %d cars\n", searchResult.TotalHits())

	// 遍历结果
	var t map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(t)) { // 适当修改以匹配你的数据结构
		if t, ok := item.(map[string]interface{}); ok {
			fmt.Printf("Found: %v\n", t)
		}
	}
}
