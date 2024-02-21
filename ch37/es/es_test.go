package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"testing"
)

func TestEs(t *testing.T) {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), // 使用你的Elasticsearch地址
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	//query := elastic.NewBoolQuery()
	// 使用NewTermsQuery构建查询
	vinTermsQuery := elastic.NewTermsQuery("vin.keyword", "HLX33B125P1184429", "HLX33B125P1186666")

	//query = query.Filter(vinTermsQuery)
	// 执行查询
	searchResult, err := client.Search().
		Index("your_index_name"). // 替换为你的索引名
		Query(vinTermsQuery).
		//Query(query).
		Sort("report_time", true). // 根据需要进行排序
		From(0).Size(10).          // 分页参数
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	// 打印查询结果
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found a total of %d documents\n", searchResult.TotalHits())

	// 遍历结果
	var mp map[string]interface{}
	for _, item := range searchResult.Each(reflect.TypeOf(mp)) {
		if obj, ok := item.(map[string]interface{}); ok {
			fmt.Printf("Found Document: %v\n", obj)
		}
	}
}

func TestIn(t *testing.T) {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), // 使用你的Elasticsearch地址
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 定义要插入的文档列表
	docs := []map[string]interface{}{
		{
			"vin":         "HLX33B125P1189999",
			"car_series":  "X01",
			"version":     "01B4",
			"message":     "2024-02-02 09:45:55.103  mcu sys I (ctrl_rx_cb:290) : CH -AEF",
			"url":         "https://ssp-uploader-biz-service.prod.k8s.chehejia.com/files/download?fileKey=vehicle/com.lixiang.datacenter/X01/20240202/HLX33B125P1184429/data/http/log/XCU/log/MCU/XCU_MCU-18442900001010360-2024_02_02_09_45_57.zip",
			"report_time": "2024-02-02 09:45:50",
			"car_usage":   "销售用车",
			"car_status":  "已激活",
			"map_id":      "2126",
			"line_no":     377,
		},
		{
			"vin":         "HLX33B125P1185555",
			"car_series":  "X01",
			"version":     "01B4",
			"message":     "2024-02-02 19:45:50.103  mcu sys I (ctrl_rx_cb:290) : Indexed document mMarzI0BeS3sAAyP2QWu to index your_index_name",
			"url":         "https://ssp-uploader-biz-service.prod.k8s.chehejia.com/files/download?fileKey=vehicle/com.lixiang.datacenter/X01/20240202/HLX33B125P1184429/data/http/log/XCU/log/MCU/XCU_MCU-18442900001010360-2024_02_02_09_45_57.zip",
			"report_time": "2024-02-02 09:45:50",
			"car_usage":   "销售用车",
			"car_status":  "已激活",
			"map_id":      "2126",
			"line_no":     377,
		},
		// 添加更多文档到列表...
	}

	// 创建批量请求
	bulkRequest := client.Bulk()

	// 为每个文档添加一个批量插入操作
	for _, doc := range docs {
		indexReq := elastic.NewBulkIndexRequest().Index("your_index_name").Doc(doc) // 替换为你的索引名称
		bulkRequest = bulkRequest.Add(indexReq)
	}

	// 执行批量操作
	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		log.Fatalf("Failed to execute bulk operation: %s", err)
	}

	// 检查是否有错误发生
	if bulkResponse != nil && bulkResponse.Errors {
		for _, item := range bulkResponse.Failed() {
			fmt.Printf("Error: %s: %s\n", item.Id, item.Error.Reason)
		}
	} else {
		fmt.Println("Bulk operation succeeded")
	}
}
