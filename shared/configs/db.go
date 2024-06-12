package configs

import (
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db = gorm.DB

var db *Db

func GetDb() *Db {
	env := GetEnv()
	if db == nil {
		var err error = nil
		db, err = gorm.Open(postgres.Open(env.DatabaseUrl), &gorm.Config{})
		utils.AssertErr(err)
	}
	return db
}
