package cmd

import (
	"github.com/urfave/cli"
	"go-admin-svr/internal/log"
	"go-admin-svr/internal/route"
)

var Admin = cli.Command{
	Name:        "admin",
	Usage:       "启动admin服务",
	Description: `admin应用服务接口`,
	Action:      runAdmin,
	Flags: []cli.Flag{
		stringFlag("port, p", "8801", "端口: 8801"),
		stringFlag("cfg, c", "app.yaml", "配置: app.yaml"),
		stringFlag("log, l", "admin.log", "日志: admin.log"),
	},
}

func runAdmin(c *cli.Context) error {

	err := GlobalInit(c.String("cfg"), c.String("log"))
	if err != nil {
		panic("Failed to initialize application: ")
		return err
	}

	engine := route.InitAdmin()
	if engine == nil {
		log.Zap.Errorw("runAdmin", "Route Engine", err)
		return err
	}

	log.Zap.Infow("runAdmin", "Starting Port", c.String("port"))
	return engine.Run(":" + c.String("port"))
}
