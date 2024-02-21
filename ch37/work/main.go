package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Vehicle struct {
	ID  uint   `gorm:"primaryKey"`
	VIN string `gorm:"column:vin;type:varchar(255);index"`
}

var db *gorm.DB
var (
	host     = pflag.StringP("host", "H", "localhost:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "123456", "Password for access to mysql service")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// 自动迁移，仅为示例，实际情况可能不需要
	db.AutoMigrate(&Vehicle{})

	r := gin.Default()
	r.GET("/vehicles", getUniqueVIN)
	r.Run(":8080")
}

func getUniqueVIN(c *gin.Context) {
	var vehicles []Vehicle
	var result []string

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	vinQuery := c.DefaultQuery("vin", "")

	query := db.Model(&Vehicle{}).Select("DISTINCT(vin)").Where("vin LIKE ?", "%"+vinQuery+"%").Offset(page * pageSize).Limit(pageSize).Find(&vehicles)

	for _, vehicle := range vehicles {
		result = append(result, vehicle.VIN)
	}

	if query.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve VINs"})
		return
	}

	c.JSON(200, result)
}
