package cmd

import (
	"github.com/urfave/cli"
	"go-admin-svr/internal/log"
	"go-admin-svr/internal/route"
)

var Api = cli.Command{
	Name:        "api",
	Usage:       "启动API服务",
	Description: `API应用接口服务`,
	Action:      runApi,
	Flags: []cli.Flag{
		stringFlag("port, p", "8802", "端口: 8802"),
		stringFlag("cfg, c", "app.yaml", "配置: app.yaml"),
		stringFlag("log, l", "api.log", "日志: api.log"),
	},
}

func runApi(c *cli.Context) error {
	var err error
	err = GlobalInit(c.String("cfg"), c.String("log"))
	if err != nil {
		panic("Failed to initialize application: ")
		return err
	}

	engine := route.InitApi()
	if engine == nil {
		log.Zap.Errorw("runApi", "Route Engine", err)
		return err
	}

	log.Zap.Infow("runApi", "Starting Port:", c.String("port"))
	return engine.Run(":" + c.String("port"))
}
