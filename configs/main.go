package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	LogLevel   int    `yaml:"log_level"`
	LogFile    string `yaml:"log_file_path"`
	Host       string `yaml:"host"`
	HttpScheme string `yaml:"http_scheme"`
	VerifySSL  bool   `yaml:"verify_ssl"` // defaults to false if blank
	AWSRegion  string `yaml:"aws_region"`
	Service    string `yaml:"service"` // "es"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readYaml() Settings {
	s := Settings{}

	yamlFile := "settings.yml"
	data, err := ioutil.ReadFile(yamlFile)
	checkErr(err)

	err = yaml.Unmarshal([]byte(data), &s)
	checkErr(err)
	return s
}

var s Settings = readYaml()
var LogLevel = s.LogLevel
var LogFile = s.LogFile
var Host = s.Host
var HttpScheme = s.HttpScheme
var VerifySSL = s.VerifySSL
var AWSRegion = s.AWSRegion
var Service = s.Service

func init() { // initializing and defaulting yaml settings
	if LogLevel == 0 {
		LogLevel = 30 // defaults to WARN level
	}

	if HttpScheme != "https" {
		HttpScheme = "http"
	}

	if AWSRegion == "" {
		panic("aws_region must be given in settings.yml")
	}
}
