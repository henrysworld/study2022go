package main

import (
	"bytes"
	"log"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dictionary struct {
	ID                       int    `gorm:"primaryKey" json:"id"`
	FirstLevelMenu           string `gorm:"type:varchar(255)" json:"first_level_menu"`
	SecondLevelMenu          string `gorm:"type:varchar(255)" json:"second_level_menu"`
	ThirdLevelMenu           string `gorm:"type:varchar(255)" json:"third_level_menu"`
	ButtonName               string `gorm:"type:varchar(255)" json:"button_name"`
	ButtonDescription        string `gorm:"type:text" json:"button_description"`
	EventType                string `gorm:"type:varchar(255)" json:"event_type"`
	RouteAddress             string `gorm:"type:varchar(255)" json:"route_address"`
	TriggerBounce            int    `gorm:"type:int" json:"trigger_bounce"`
	HasPermissionRestriction int    `gorm:"type:int" json:"has_permission_restriction"`
	OperationPermission      string `gorm:"type:varchar(255)" json:"operation_permission"`
	IsExternalUsage          int    `gorm:"type:int" json:"is_external_usage"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Dictionary{})

	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the uploaded file"})
			return
		}

		buf := bytes.NewBuffer(nil)
		_, err = buf.ReadFrom(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file buffer"})
			return
		}

		f, err := excelize.OpenReader(buf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read excel file"})
			return
		}

		rows, err := f.GetRows("Sheet1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rows from excel sheet"})
			return
		}
		for _, row := range rows[1:] {
			dict := Dictionary{
				FirstLevelMenu:           row[0],
				SecondLevelMenu:          row[1],
				ThirdLevelMenu:           row[2],
				ButtonName:               row[3],
				ButtonDescription:        row[4],
				EventType:                row[5],
				RouteAddress:             row[6],
				TriggerBounce:            convertToBoolInt(row[7]),
				HasPermissionRestriction: convertToBoolInt(row[8]),
				OperationPermission:      row[9],
				IsExternalUsage:          convertToBoolInt(row[10]),
			}

			db.Create(&dict)
		}

		c.JSON(http.StatusOK, gin.H{"status": "Data imported successfully"})
	})

	// 按条件查询接口
	// 按条件查询接口
	r.GET("/dictionary", func(c *gin.Context) {
		dicts, _ := FetchDictionaryByConditions(&Dictionary{}, db)

		c.JSON(http.StatusOK, dicts)
	})

	r.Run(":8080")
}
func convertToBoolInt(value string) int {
	if value == "是" {
		return 1
	}
	return 0
}

// 将驼峰命名法转换为下划线分隔（为了匹配数据库中的字段名）
func snakeString(s string) string {
	var result string
	var lastUpper bool
	for _, v := range s {
		if unicode.IsUpper(v) {
			if !lastUpper {
				result += "_"
			}
			lastUpper = true
		} else {
			lastUpper = false
		}
		result += string(unicode.ToLower(v))
	}
	return result
}

// FetchDictionaryByConditions fetches Dictionary items based on given conditions
func FetchDictionaryByConditions(conditions *Dictionary, db *gorm.DB) ([]Dictionary, error) {
	var result []Dictionary
	query := db
	t := reflect.TypeOf(*conditions)
	v := reflect.ValueOf(*conditions)

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("gorm")
		if strings.Contains(columnName, "column:") {
			columnName = strings.Split(columnName, ":")[1]
			if value := v.Field(i).Interface(); value != "" {
				query = query.Where(columnName+" = ?", value)
			}
		}
	}

	err := query.Find(&result).Error
	return result, err
}
