package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"goProject/student/server/config"
	"time"
)

// 创建NewClubDBConn连接
func NewStudentDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %s no set.", dbName))
	}

	studentsDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%s", err.Error()))
	}
	// 设置最大链接数
	studentsDB.DB().SetMaxOpenConns(10)
	studentsDB.DB().SetMaxIdleConns(2)
	studentsDB.DB().SetConnMaxLifetime(3 * time.Minute)
}
