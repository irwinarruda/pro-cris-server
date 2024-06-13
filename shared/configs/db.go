package configs

import (
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db = *gorm.DB

var db Db

func GetDb() Db {
	if db == nil {
		env := GetEnv()
		var err error = nil
		db, err = gorm.Open(postgres.Open(env.DatabaseDsn), &gorm.Config{})
		utils.AssertErr(err)
	}
	return db
}
