//author: richard
package server

import (
	"database/sql"
	"fmt"
	lm "github.com/advancevillage/letaotao/manager"
	"sync"
)

var conn *sql.DB
var err	 error
var config *lm.DatabaseConfig
var driver string
var once sync.Once

func InitDatabase(d string , c *lm.DatabaseConfig) error {
 	once.Do(func() {
 		driver = d
 		config = c
 		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", config.User, config.Password, config.Host, config.Port, config.Schema, config.CharSet)
		conn, err = sql.Open(driver, dns)
	})
 	return err
}

func Conn() *sql.DB {
	once.Do(func() {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", config.User, config.Password, config.Host, config.Port, config.Schema, config.CharSet)
		conn, err = sql.Open(driver, dns)
	})
	return conn
}


