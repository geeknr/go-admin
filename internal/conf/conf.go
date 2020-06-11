package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"go-admin-svr/internal/util/pathutil"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Init(customConf string) error {
	var err error
	if customConf == "" {
		customConf = filepath.Join(pathutil.WorkDir(), "conf", "app.yaml")
	} else {
		customConf = filepath.Join(pathutil.WorkDir(), "conf", customConf)
	}

	viper.SetConfigFile(customConf)
	content, err := ioutil.ReadFile(customConf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	return nil
}
