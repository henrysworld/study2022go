package main

import (
	//"encoding/json"
	"fmt"
	"gorm.io/gorm/logger"
	"net"
	"os"
	"strings"
	"time"

	"github.com/henrysworld/study2022go/ch37/pkg/log"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	ID        uint64    `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Code      string    `gorm:"column:code"`
	Price     uint      `gorm:"column:price"`
	CreatedAt time.Time `json:"createAt,omitempty" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	DeletedAt time.Time `json:"-" gorm:"column:deletedAt;index:idx_deletedAt"`
	FileMd5   string    `json:"fileMd5" gorm:"column:fileMd5;index:idx_fileMd5"`
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deletedAt"`
}

type ProductRes struct {
	// gorm.Model
	ID    uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
	// DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deletedAt"`
}

func (p *Product) TableName() string {
	return "product"
}

var (
	host     = pflag.StringP("host", "H", "localhost:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "123456", "Password for access to mysql service")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func main() {
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// logger 配置
	opts := &log.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true,
		DisableCaller:    true,
		OutputPaths:      []string{"test.log", "stdout"},
		ErrorOutputPaths: []string{"error.log"},
	}

	//初始化全局Logger
	log.Init(opts)
	defer log.Flush()

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(4),
	})
	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()

	_, ok := tx.InstanceGet("gorm:started_transaction")
	host, _ := os.Hostname()
	addr := InternalIP()
	log.Infof("Host: %s", host)
	log.Infof("addr: %s", addr)
	if ok {
		log.Infof("处于事务中")
	} else {
		log.Infof("不处于事务中")
	}

	//db.AutoMigrate(&Product{})

	if err := db.Table("product_xcu").Create(&Product{Code: "D76", Price: 89, FileMd5: "zzzzzzzzz", DeletedAt: time.Now()}).Error; err != nil {
		log.Fatalf("Create error: %v", err)
	}

	//pluckRecords(db)
	//updateRecords(db)

	//distinctOrderPluckRecords(db)
	//testPb()
	getCs(db)
	//deleteRecords(db)
	//getCs(db)

	//for i := 0; i < 100; i++ {
	//	go insertC(db)
	//}

	//ch := make(chan int, 1)
	//
	//<-ch

	//db.Transaction(func(tx *gorm.DB) error {
	//	err := tx.Create(&UserM{
	//		Username: "henry",
	//		Nickname: "henry",
	//		Password: "henry",
	//		Phone:    "17612292240",
	//		Email:    "442168402@qq.com",
	//	}).Error
	//	err = fmt.Errorf("err: test")
	//	if err != nil {
	//		log.Errorf("err: ", err.Error())
	//		return err
	//	}
	//
	//	return nil
	//})

	// d := p.db.Exec("delete from policy_audit where deletedAt < ?", date)

	//PrintProducts(db)

	product := &Product{}
	if err := db.Where("code = ?", "D42").First(&product).Error; err != nil {
		log.Fatalf("Create product error: %v", err)
	}

	// product.Price = 200
	// if err := db.Save(product).Error; err != nil {
	// 	log.Fatalf("Update product error: %v", err)
	// }

	// PrintProducts(db)

	// date := time.Now().AddDate(0, 0, -maxReserveDays).Format("2006-01-02 15:04:05")

	// d := p.db.Exec("delete from policy_audit where deletedAt < ?", date)

	// if err := db.Where("code = ?", "D42").Delete(&Product{}).Error; err != nil {
	// 	log.Fatalf("Delete product error: %v", err)
	// }

	// PrintProducts(db)

	// fmt.Println(json.Marshal(product))

}

func testPb() {
	//meminfoField := proto.MemInfoField{Mem: []int32{7777}}
	//fs := make([]*proto.MemInfoField, 0)
	//fs = append(fs, &meminfoField)
	//msg := &proto.ContinueMsgResponse{
	//	Msg:          "aaa",
	//	MemInfoField: fs,
	//}
	//
	//m := jsonpb.Marshaler{
	//	OrigName:     false,
	//	EnumsAsInts:  false,
	//	EmitDefaults: false,
	//	Indent:       "",
	//	AnyResolver:  nil,
	//}
	//data, _ := m.MarshalToString(msg)
	//
	//jsonStr := string(data)
	//fmt.Printf("%s", jsonStr)

	//{"msg":"aaa","memInfoField":[[7777]]}

}

func insertC(db *gorm.DB) {

	if err := db.Create(&Product{Code: "D42", Price: 100}).Error; err != nil {
		log.Fatalf("Create error: %v", err)
	}
}
func getCs(db *gorm.DB) {
	var products []Product
	ids := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var ss []string
	//for _, id := range ids {
	//	ss = append(ss, fmt.Sprintf("%d", id))
	//}
	//s := strings.Join(ss, ",")
	err := db.Raw("select * from product where id in (?) limit ?", ids, 2).Scan(&products).Error
	if err != nil {
		fmt.Println(err)
	}
	var productsxcu []Product
	db.Table("product_xcu").Where("id in (?)", ids).Limit(2).Find(&productsxcu)
	//fmt.Println(products)
	//p, _ := json.Marshal(products)
	fmt.Printf("1=====%+v\n", products)
	fmt.Printf("2=====%+v\n", productsxcu)
	//db.Row("select * from product where id in (?)", ids...).Scan(products)
}

type User struct {
	IDS int64 `json:"ids"`
}

func deleteRecords(db *gorm.DB) {
	//var products = []User{{ID: 1}, {ID: 2}, {ID: 3}}
	//db.Table("product").Delete(&products)
}

func pluckRecords(db *gorm.DB) {
	sql := `select * from product group by code;`
	var ids []string
	db.Raw(sql).Pluck("code", &ids)
	fmt.Printf("ids: %v", ids)
}

func updateRecords(db *gorm.DB) {
	sql := `update product set price = ? where code = "D42";`
	var args []interface{}
	args = append(args, "666")
	db.Raw(sql, args)
	db.Exec(sql, args)
}

func distinctOrderPluckRecords(db *gorm.DB) {
	//var codes []*string
	//if err := db.Table("product").Select("distinct(code), createdAt").Where("code=? and price = ?", "D42", 666).Order("createdAt").Pluck("code", &codes); err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("ch==%s", codes)

	//var codes []string
	var cond []any
	cond = append(cond, "D42")
	cond = append(cond, 666)
	var ps []Product
	if err := db.Table("product").
		Select("DISTINCT(code) as code, createdAt").
		Where("code = ? AND price = ?", cond).
		Order("createdAt").
		Scan(&ps).
		//Pluck("createdAt", &codes).
		Error; err != nil {
		fmt.Println(err)
	}

	fmt.Printf("codes: %+v", ps)

}

func updateC(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		e := tx.Exec("UPDATE product SET price = price + 1 WHERE id = ?", 1).RowsAffected
		fmt.Println(e)

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func PrintProducts(db *gorm.DB) {
	products := make([]*Product, 0)
	var count int64
	d := db.Unscoped().Where("code like ?", "%D%").Offset(0).Limit(-1).Order("id desc").Find(&products).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error: %v", d.Error)
	}

	log.V(log.DebugLevel).Infof("totalcount: %d", count)
	for _, product := range products {
		log.V(log.DebugLevel).Infof("code: %s, price: %d", product.Code, product.Price)
	}
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("CH==========BeforeCreate")
	tx.Statement.Table = "product_1"
	//tx = tx.Table("product_1")
	return
}
func (u *Product) BeforeFind(tx *gorm.DB) (err error) {
	fmt.Println("CH==========BeforeFind")
	//tx.Statement.Table = GetUserTable(tx.Statement.Context, u.TableName())
	return
}
