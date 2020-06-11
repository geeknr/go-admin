package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-admin-svr/internal/conf"
	"go-admin-svr/internal/db"
	"go-admin-svr/internal/log"
	"go-admin-svr/internal/util/pathutil"
	"go-admin-svr/internal/util/sonyflake"
)

func GlobalInit(customConf, customLog string) error {
	var err error
	err = conf.Init(customConf)
	if err != nil {
		return errors.Wrap(err, "init configuration")
	}

	err = log.Init(customLog)
	if err != nil {
		log.Zap.Errorw("Failed to initialize Logging: %v", err)
		return err
	}

	if err := db.Init(); err != nil {
		log.Zap.Errorw("adminserver", "Failed to initialize Mysql cache", err)
		return err
	}

	sonyflake.IdCter.Init(viper.GetInt("serverID"))

	log.Zap.Infow("GlobalInit", "APP", viper.GetString("app"), "VERSION", viper.GetString("version"))
	log.Zap.Infow("GlobalInit", "Work directory", pathutil.WorkDir())

	return nil
}
