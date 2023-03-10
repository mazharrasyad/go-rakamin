package container

import (
	"fmt"
	"os"
	"path/filepath"

	"tugas_akhir/internal/infrastructure/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

type (
	Container struct {
		Mysqldb *gorm.DB
		Apps    *Apps
	}

	Apps struct {
		Name      string `mapstructure:"name"`
		Version   string `mapstructure:"version"`
		Address   string `mapstructure:"address"`
		HttpPort  int    `mapstructure:"httpPort"`
		SecretJwt string `mapstructure:"secretJwt"`
	}
)

func init() {
	v = viper.New()

	v.AutomaticEnv()
	v.SetConfigFile(".env")
	// v.AddConfigPath("./.env")

	path, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("failed read config : %s", err.Error()))
	}

	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("failed init config : %s", err.Error()))
	}
}

func AppsInit(v *viper.Viper) (apps Apps) {

	err := v.Unmarshal(&apps)
	if err != nil {
		panic(fmt.Sprintf("failed init database mysql : %s", err.Error()))
	}
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	mysqldb := mysql.DatabaseInit(v)

	return &Container{
		Apps:    &apps,
		Mysqldb: mysqldb,
	}

}
