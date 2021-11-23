package logic

import (
	"github.com/belief428/gorm-engine/tools"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	Path string
	Name string
}

func (this *Sqlite) DSN() gorm.Dialector {
	if isExist, _ := tools.PathExists(this.Path); !isExist {
		_ = tools.MkdirAll(this.Path)
	}
	return sqlite.Open(this.Path + "/" + this.Name)
}
