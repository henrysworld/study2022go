package tch

//
//import (
//	"fmt"
//	pb "github.com/henrysworld/study2022go/ch37/grpc/proto"
//	"gorm.io/gorm"
//	"time"
//
//	//v2 "github.com/henrysworld/study2022go/ch37/grpc/proto/v2"
//	"testing"
//)
//
//func TestName(t *testing.T) {
//	userInfo := pb.UserInfo{
//		Username:  "",
//		Nickname:  "",
//		Email:     "",
//		Phone:     "",
//		PostCount: 0,
//	}
//
//	userInfoV2 := v2.UserInfo{
//		Username:    "",
//		Nickname:    "",
//		Email:       "",
//		Phone:       "",
//		PostCountV2: 0,
//	}
//
//	fmt.Printf("%v\n", userInfo)
//	fmt.Printf("%v\n", userInfoV2)
//}
//
//type Base struct {
//	ID        uint `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time `sql:"index"`
//}
//
//type User struct {
//	Base
//	Username string
//	Email    string
//}
//
//func GetUserTable(userID uint, baseTableName string) string {
//	return fmt.Sprintf("%s_%d", baseTableName, userID%100)
//}
//
//func CreateUser(db *gorm.DB, user *User) error {
//	tableName := GetUserTable(user.ID, "users")
//	return db.Table(tableName).Create(user).Error
//}
//
//func GetUser(db *gorm.DB, userID uint) (*User, error) {
//	var user User
//	tableName := GetUserTable(userID, "users")
//	if err := db.Table(tableName).Where("id = ?", userID).First(&user).Error; err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//type IUser interface {
//	GetUser1(db *gorm.DB, userID uint) (*User, error)
//}
//type user struct {
//	db *gorm.DB
//}
//
//func (u *user) GetUser1(db *gorm.DB, userID uint) (*User, error) {
//	var user User
//	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
