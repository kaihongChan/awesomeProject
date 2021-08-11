package initialize

import (
	"awesomeProject/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormInit() *gorm.DB {
	switch global.Config.GormConf.Driver {
	case "mysql":
		return GormMysql()
	// code here...
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.Config.MysqlConf
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Params
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConn)
		sqlDB.SetMaxOpenConns(m.MaxOpenConn)
		return db
	}
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}
