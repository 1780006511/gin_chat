package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs/")
	viper.SetConfigType("yml") //设置配置文件类型
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("configs app:", viper.Get("app"))

}

func InitMySQL() {
	//自定义日志打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      false,       //彩色
		})
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})

	fmt.Println("configs mysql:", viper.Get("mysql"))
	//user := models.UserBasic{}
	//
	//DB.Find(&user)
	//fmt.Println(user)
}
