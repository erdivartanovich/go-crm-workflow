package setting

import (
	"fmt"
	"path/filepath"

	ini "gopkg.in/ini.v1"
)

type DatabaseConfig struct {
	Host     string `ini:"HOST"`
	Name     string `ini:"NAME"`
	User     string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Driver   string `ini:"DRIVER"`
	Charset  string `ini:"CHARSET"`
	Debug    bool   `ini:"DEBUG"`
}

var (
	//ConfigFile  configuration file path
	ConfigFile string = "./config.ini.example"
	Config            = &struct {
		ApiHost    string `ini:"API_HOST"`
		ApiPort    string `ini:"API_PORT"`
		TimeFormat string `ini:"TIME_FORMAT"`
	}{
		ApiPort:    ":8080",
		TimeFormat: "2006-01-02 15:04:05",
	}
	Db = &DatabaseConfig{}
)

func (db DatabaseConfig) GetDataSourceName() string {

	switch db.Driver {
	case "mysql":
		if db.Host[0] == '/' {
			return fmt.Sprintf(
				"%s:%s@unix(%s)/%s?charset=%s&parseTime=true",
				db.User,
				db.Password,
				db.Host,
				db.Name,
				db.Charset,
			)
		}
		return fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
			db.User,
			db.Password,
			db.Host,
			db.Name,
			db.Charset,
		)
	}
	return ""
}

//Initialize Initialize configuration
func Initialize() error {

	path, err := BasePath()

	if err != nil {
		return err
	}

	path = path + "/" + ConfigFile

	cfg, err := ini.Load(path)
	err = cfg.Section("application").MapTo(Config)
	if err != nil {
		return err
	}

	err = cfg.Section("database").MapTo(Db)

	if err != nil {
		return err
	}

	return nil
}

func BasePath() (string, error) {
	path, err := filepath.Abs("./")

	return path, err
}
