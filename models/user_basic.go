package models

import (
	"fmt"
	"gin_chat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"name"`
	PassWord      string    `gorm:"pass_word"`
	Phone         string    `gorm:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string    `gorm:"email" valid:"email"`
	Identity      string    `gorm:"identity"`
	ClientIp      string    `gorm:"client_ip"`
	ClientPort    string    `gorm:"client_port"`
	Salt          string    `gorm:"salt"`
	LoginTime     time.Time `gorm:"login_time"`
	HeartbeatTime time.Time `gorm:"heartbeat_time"`
	LogoutTime    time.Time `gorm:"logout_time"`
	IsLogout      bool      `gorm:"is_logout"`
	DeviceInfo    string    `gorm:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	//for _, v := range data {
	//	fmt.Println()
	//	fmt.Println(v)
	//}
	return data
}

func FindUserByNameAndPassword(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.EncodeMd5(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) (user UserBasic, err error) {
	user = UserBasic{}
	err = utils.DB.Where("name = ?", name).First(&user).Error
	return
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email})
}
