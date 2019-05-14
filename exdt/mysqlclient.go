package exdt

import (
	"github.com/gohouse/gorose"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type MySQLClient struct {
}

var conn *gorose.Connection
var mc sync.Once

func (c *MySQLClient) Instance(dns, prefix string) *gorose.Connection {
	mc.Do(func() {
		var err error
		conn, err = gorose.Open(getDbConfig(dns, prefix))
		if err != nil {
			panic(err)
		}
	})
	return conn
}

/**
获取数据库的配置,后期放到配置文件进行管理
 */
func getDbConfig(dns, prefix string) (*gorose.DbConfigSingle) {
	return &gorose.DbConfigSingle{
		Driver:          "mysql", // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,    // 是否开启sql日志
		SetMaxOpenConns: 0,       // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,       // (连接池)闲置的连接数
		Prefix:          prefix,  // 表前缀
		Dsn:             dns,     // 数据库链接
	}
}
