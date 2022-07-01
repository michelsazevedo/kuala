package config

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Secret string `yaml:"secret"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
}

//Config ...
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		Username string `yaml:username`
		Password string `yaml:password`
		Timeout  string `yaml:timeout`
	} `yaml:"database"`

	Settings Settings `yaml:"settings"`
}

//NewConfig ...
func NewConfig(configs string) (*Config, error) {
	file, err := ExpandEnv(configs)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ExpandEnv(configs string) (io.Reader, error) {
	file, err := os.Open(configs)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bufferConfigs := new(bytes.Buffer)
	_, err = bufferConfigs.ReadFrom(file)
	if err != nil {
		return nil, err
	}

	bytesConfigs := []byte(os.ExpandEnv(bufferConfigs.String()))
	return bytes.NewReader(bytesConfigs), nil
}
