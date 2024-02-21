package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()
	//rows, err := conn.Query(ctx, "SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 5")
	rows, err := conn.Query(ctx, "select count(*) count from system.clusters group by shard_num")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			//name, uuid string
			count uint64
		)
		if err := rows.Scan(
			//&name,
			//&uuid,
			&count,
		); err != nil {
			log.Fatal(err)
		}
		//log.Printf("name: %s, uuid: %s",
		//	name, uuid)
		fmt.Println(count)
	}

}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			//Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
			//Addr: []string{"tcp://192.168.0.101:9000"},
			Addr: []string{"192.168.0.101:50288", "192.168.0.101:50291", "192.168.0.101:50286", "192.168.0.101:50285"},
			Auth: clickhouse.Auth{
				Database: "system",
				Username: "default",
				Password: "123456",
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			//TLS: &tls.Config{
			//	InsecureSkipVerify: true,
			//},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}
