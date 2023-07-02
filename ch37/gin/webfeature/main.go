package main

import (
	"fmt"
	"github.com/henrysworld/study2022go/ch37/zyjgrpc"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/henrysworld/study2022go/ch37/pkg/log"
	"golang.org/x/sync/errgroup"
)

var (
	host     = pflag.StringP("host", "H", "localhost:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "sse", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "StrongeneDB123456!", "Password for access to mysql service")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
	S        *gorm.DB
)

type Product struct {
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Price       int       `json:"price" binding:"gte=0"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	Banji       Banji     `json:"banji" binding:"required"`
}

type Banji struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

var ch *zyjgrpc.Channel

func newProductHandler() *productHandler {
	return &productHandler{
		products: make(map[string]Product),
	}
}

func (p *productHandler) Create(ctx *gin.Context) {
	p.Lock()
	defer p.Unlock()
	//1.参数解析
	var product Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//2.参数校验
	if _, ok := p.products[product.Name]; ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product %s already exist", product.Name)})
		return
	}

	product.CreatedAt = time.Now()

	//3.逻辑处理
	p.products[product.Name] = product
	//log.Printf("Register product %s success", product.Name)

	//4.返回结果
	ctx.JSON(http.StatusOK, product)

}

func (p *productHandler) Get(ctx *gin.Context) {

	ch := zyjgrpc.NewChannel()
	ch.Push(ctx.Param("name"))
	//arrsy := getCs(S)
	//
	//fmt.Println(arrsy)
	//
	//p.RLock()
	//defer p.RUnlock()

	//product, ok := p.products[ctx.Param("name")]
	//if !ok {
	//	ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("can not found product %s", ctx.Param("name"))})
	//	return
	//}

	ctx.JSON(http.StatusOK, ctx.Param("name"))
}

func getCs(db *gorm.DB) []Product {
	var products []Product
	ids := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var ss []string
	//for _, id := range ids {
	//	ss = append(ss, fmt.Sprintf("%d", id))
	//}
	//s := strings.Join(ss, ",")
	err := db.Raw("select * from product where id in (?)", ids).Scan(&products).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(products)
	//p, _ := json.Marshal(products)
	fmt.Printf("%+v", products)
	//db.Row("select * from product where id in (?)", ids...).Scan(products)

	return products
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()
	//路由分组、中间件、认证
	v1 := router.Group("/v1")
	{
		productv1 := v1.Group("/products")
		{
			//路由匹配
			productv1.POST("", productHandler.Create)
			productv1.GET(":name", productHandler.Get)
		}
	}

	return router
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

	ch = zyjgrpc.NewChannel()

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
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(4),
	})

	S = db

	var eg errgroup.Group

	//一进行多端口
	insecureServer := &http.Server{
		Addr:         ":8080",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// secureServer := &http.Server{
	// 	Addr:         ":8443",
	// 	Handler:      router(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			//log.Fatal(err)
		}

		return err
	})

	// eg.Go(func() error {
	// 	err := secureServer.ListenAndServeTLS("server.pem", "server.key")
	// 	if err != nil && err != http.ErrServerClosed {
	// 		log.Fatal(err)
	// 	}
	// 	return err
	// })

	if err := eg.Wait(); err != nil {
		//log.Fatal(err)
	}

}
