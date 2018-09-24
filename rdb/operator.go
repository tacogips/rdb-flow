package rdb

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func DropAllTable(conf ConnectionConfig) error {
	conn, err := newConn(conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	tables, err := conn.DBMetas()
	if err != nil {
		return err
	}

	err = conn.DropTables(tables)
	if err != nil {
		return err
	}

	return nil
}

func Execute(conf ConnectionConfig, sql string, args ...interface{}) error {
	conn, err := newConn(conf)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func mysqlConnStringFromConfig(conf ConnectionConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/db?parseTime=true")
}

func newConn(conf ConnectionConfig) (*xorm.Engine, error) {

	switch conf.Driver {
	case "mysql":
		return xorm.NewEngine("mysql", mysqlConnStringFromConfig(conf))
	default:
		panic("not implemented driver %s", conf.Driver)
	}
}
