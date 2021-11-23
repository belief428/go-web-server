package logic

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	User, Password, Host string
	Port                 int
	DBName, Parameters   string
}

func (this *Mysql) DSN() gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		this.User, this.Password, this.Host, this.Port, this.DBName, this.Parameters,
	))
}
