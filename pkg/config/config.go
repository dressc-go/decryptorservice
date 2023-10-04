package config

import (
	"github.com/dressc-go/zlogger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	ConfigFile        string
	IpAddress         string `yaml:"IpAddress"`
	Port              uint16 `yaml:"Port"`
	EncPubKeyFile     string `yaml:"EncPubKeyFile"`
	DecPrivateKeyFile string `yaml:"DecPrivateKeyFile"`
	TLSCertFile       string `yaml:"TLSCertFile"`
	TLSKeyFile        string `yaml:"TLSKeyFile"`
}

func (cnf *Config) New() error {
	logger := zlogger.GetLogger("Config.New")

	cnf.ConfigFile = getConfFilePath()
	if cnf.ConfigFile == "" {
		e := errors.New("Config file does not exist (or was not given in ENV: ConfigFile")
		logger.Error().Err(e).Msg("")
		return e
	}
	yamlFile, err := ioutil.ReadFile(cnf.ConfigFile)
	if err != nil {
		e := errors.Wrap(err, "Could not open config file:"+cnf.ConfigFile)
		logger.Error().Err(e).Msg("")
		return e
	}
	err = yaml.UnmarshalStrict(yamlFile, &cnf)
	if err != nil {
		e := errors.Wrap(err, "Could not read config file:"+cnf.ConfigFile)
		logger.Error().Err(e).Msg("")
		return e
	}
	return nil
}

func getConfFilePath() string {
	logger := zlogger.GetLogger("Config.getConfFilePath")

	explicitFile := os.Getenv("ConfigFile")
	if explicitFile != "" {
		if fileExists(explicitFile) {
			logger.Debug().Msg("Config selected from ENV Configfile: " + explicitFile)
			return explicitFile
		} else {
			e := errors.New("ENV ConfigFile " + explicitFile + "was given, but not found")
			logger.Error().Err(e).Msg("")
			return ""
		}
	}

	var candidates [2]string
	candidates[0] = "config.yml"
	candidates[1] = "/run/secrets/config.yml"

	for _, candidate := range candidates {
		if fileExists(candidate) {
			logger.Debug().Msg("Config selected: " + candidate)
			return candidate
		}
		logger.Debug().Msg("Tried configfile: " + candidate + ": did not exist")
	}
	return ""
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
