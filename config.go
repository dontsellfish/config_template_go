package config_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
)

type Config struct {
	Example string `json:"example"`
	Sample  int    `json:"sample"`
	Sub     struct {
		SubData int `json:"sub_data"`
	} `json:"sub"`
}

func LoadConfig(filename string) (*Config, error) {
	buff, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (cfg Config) DumpConfig(filename string) error {
	buffer, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, buffer, os.ModePerm)
}

type ConfigUtil struct {
	Data           *Config
	Filename       string
	BackupFilename string
}

func NewConfigUtil(filename string) (*ConfigUtil, error) {
	cfg, err := LoadConfig(filename)
	if err != nil {
		return nil, err
	}
	return WrapConfigUtil(cfg, filename)
}

func WrapConfigUtil(cfg *Config, filename string) (*ConfigUtil, error) {
	dir, file := path.Split(filename)
	return &ConfigUtil{cfg, filename, path.Join(dir, fmt.Sprintf(".backup_%s", file))}, nil
}

func (util *ConfigUtil) Backup() error {
	err := util.Data.DumpConfig(util.BackupFilename)
	if err != nil {
		return err
	}
	return nil
}

func (util *ConfigUtil) Rollback() (err error) {
	util.Data, err = LoadConfig(util.BackupFilename)
	if err != nil {
		return err
	}
	return os.Rename(util.BackupFilename, util.Filename)
}

func Update(data map[string]interface{}, req ...interface{}) error {
	if len(req) < 2 {
		return nil
	} else {
		switch req[0].(type) {
		case string:
			_, ok := data[req[0].(string)]
			if !ok {
				return errors.New(fmt.Sprintf("key '%s' is not found", req[0].(string)))
			}

			if len(req) == 2 {
				data[req[0].(string)] = req[1]
				return nil
			} else {
				return Update(data[req[0].(string)].(map[string]interface{}), req[1:]...)
			}
		default:
			return errors.New(fmt.Sprintf("key '%v' isn't string", req[0]))
		}
	}
}

func (util *ConfigUtil) Update(req ...interface{}) error {
	buff, err := os.ReadFile(util.Filename)
	if err != nil {
		return err
	}
	var data map[string]interface{}
	err = json.Unmarshal(buff, &data)
	if err != nil {
		return err
	}

	err = Update(data, req...)
	if err != nil {
		return err
	}
	buff, err = json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buff, &util.Data)
	if err != nil {
		return err
	}
	err = os.Rename(util.Filename, util.BackupFilename)
	if err != nil {
		return err
	}

	return util.Data.DumpConfig(util.Filename)
}
