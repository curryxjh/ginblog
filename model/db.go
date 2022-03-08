package model

//配置数据库参数以及数据库连接
import (
	"fmt"
	"ginbolg/utils"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/mysql"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		//禁用默认事务(提高运行速度)
		SkipDefaultTransaction: true,

		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。需要注意不要超过框架时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//sqlDB.Close();
}
