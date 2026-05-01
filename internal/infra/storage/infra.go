package storage

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Infra struct {
	MySQL *gorm.DB
}
type Config struct {
	MySQL mysql.Config
}

func NewInfra(cfg Config) (*Infra, func(), error) {
	infra := &Infra{}
	mysqlDb, err := NewMySQL(cfg.MySQL)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		infra.Close()
	}
	infra.MySQL = mysqlDb
	return infra, cleanup, nil
}
func (i *Infra) Close() error {
	if i.MySQL != nil {
		return CloseMySQL(i.MySQL)
	}
	return nil
}
