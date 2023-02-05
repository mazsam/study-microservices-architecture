package demo_go_swagger_api

import (
	// "fmt"
	// "fmt"
	// "io/ioutil"
	"net/http"

	// "os"

	"github.com/go-openapi/errors"
	// "github.com/spf13/viper"
	// "gopkg.in/yaml.v3"

	"github.com/sirupsen/logrus"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"

	"gateway/configs"
	"gateway/gen/models"
)

// NewRuntime creates a new application level runtime that encapsulates the shared services for this application
func NewRuntime(cfg *configs.Config, log logrus.FieldLogger, runMigration bool) (*Runtime, error) {


	// DB
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	// cfg.DB.PostgresSQL.Username,
	// 	// cfg.DB.PostgresSQL.Password,
	// 	// cfg.DB.PostgresSQL.Host,
	// 	// cfg.DB.PostgresSQL.Port,
	// 	// cfg.DB.PostgresSQL.Name,
	// )

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	// if err != nil {
	// 	log.WithFields(logrus.Fields{
	// 		"module": "runtime",
	// 	}).Errorf("DB connection : %v", err)
	// 	os.Exit(1)
	// }

	rt := &Runtime{
		// Db:           db,
		Cfg: cfg,
		Log: log,
	}

	if runMigration {
		rt.RunMigration()
	}

	return rt, nil
}

type Runtime struct {
	Db  *gorm.DB
	Cfg *configs.Config
	Log logrus.FieldLogger
}

func (r *Runtime) DB() *gorm.DB {
	return r.Db
}

// Logger gets the root logger for this application
func (r *Runtime) Logger() logrus.FieldLogger {
	return r.Log
}

// Config returns the viper config for this application
func (r *Runtime) Config() *configs.Config {
	return r.Cfg
}

// RunMigration is run db migration
func (r *Runtime) RunMigration() {
	r.Log.Infof("Migrating DBs")
	r.Db.AutoMigrate(
		&models.Hello{},
	)
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}
