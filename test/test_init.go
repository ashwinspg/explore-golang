package test

import (
	"database/sql"
	"sync"

	"github.com/ashwinspg/explore-golang/db"
	"github.com/ashwinspg/explore-golang/utils"

	"github.com/sirupsen/logrus"
)

var once sync.Once

//Env holds a test environment required to run the tests
type Env struct {
	DBConn *sql.DB
	L      *logrus.Entry
}

//SetupTestEnv initialises and returns a complete test environment
func SetupTestEnv() Env {
	once.Do(func() {
		db.MigrateUp()
		db.InitiateConnection()
	})

	return Env{
		DBConn: db.GetDB(),
		L:      utils.LogEntryWithRef(),
	}
}
