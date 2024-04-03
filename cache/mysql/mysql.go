package mysql

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

func New() *gorm.DB {
	//建立数据库连接,数据库名:数据库密码@...
	//dsn := "root:w123456.@tcp(127.0.0.1:3306)/poem?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:niganmaaiyou@tcp(39.104.77.138:3360)/light?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//处理错误
	if err != nil {
		//控制台打印错误日志
		panic("数据库连接失败!")
	}
	return db
}
