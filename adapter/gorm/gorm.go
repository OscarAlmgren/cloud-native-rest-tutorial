package gorm

import (
	"cloud-native/config"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func New(conf *config.Conf) (*gorm.DB, error) {
	cfg := &mysql.Config{
		User:                    conf.Db.Username,
		Passwd:                  conf.Db.Password,
		Net:                     "tcp",
		Addr:                    conf.Db.Host,
		DBName:                  conf.Db.DbName,
		Params:                  map[string]string{},
		Collation:               "",
		Loc:                     &time.Location{},
		MaxAllowedPacket:        0,
		ServerPubKey:            "",
		TLSConfig:               "",
		Timeout:                 0,
		ReadTimeout:             0,
		WriteTimeout:            0,
		AllowAllFiles:           false,
		AllowCleartextPasswords: false,
		AllowNativePasswords:    true,
		AllowOldPasswords:       false,
		CheckConnLiveness:       false,
		ClientFoundRows:         false,
		ColumnsWithAlias:        false,
		InterpolateParams:       false,
		MultiStatements:         false,
		ParseTime:               true,
		RejectReadOnly:          false,
	}

	return gorm.Open("mysql", cfg.FormatDSN())
}
