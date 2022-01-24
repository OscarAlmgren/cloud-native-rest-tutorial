package db

import (
	"cloud-native/config"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func New(conf *config.Conf) (*sql.DB, error) {
	cfg := &mysql.Config{
		User:                    conf.Db.Username,
		Passwd:                  conf.Db.Password,
		Net:                     "tcp",
		Addr:                    fmt.Sprintf("%v:%v", conf.Db.Host, conf.Db.Port),
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
		ParseTime:               false,
		RejectReadOnly:          false,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}
