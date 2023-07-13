package config

import (
    "fmt"
    "io/ioutil"
    "strings"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "gopkg.in/yaml.v3"
)

var (
    Viper *viper.Viper
)

type yamlConfig struct {
    DBConfig     DBConfig     `yaml:"database"`
    ServerConfig ServerConfig `yaml:"serverDetails"`
}

type DBConfig struct {
    Server        string `json:"server"`
    Username      string `json:"username"`
    Password      string `json:"password"`
    Port          string `json:"port"`
    Schema        string `json:"schema"`
    MaxConnection int    `json:"connection_max"`
}

type ServerConfig struct {
    Port  string `json:"port"`
    Debug string `json:"debug"`
}

var cfg yamlConfig

func Initialize(path string) {
    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        panic("not able to read file")
    }
    if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
        log.Fatalln("-----> problem in unmarshalling", err)
    }
    fmt.Println(cfg.DBConfig.Schema)
}

func Get(configName string) interface{} {
    if configName == "database" {
        return cfg.DBConfig
    } else if configName == "server" {
        return cfg.ServerConfig
    }
    return nil
}

func init() {
    readConfig("config/config")
}

func readConfig(filename string) {
    Viper = viper.New()
    Viper.AddConfigPath(".")
    Viper.SetConfigName(filename)
    err := Viper.ReadInConfig()
    if err != nil {
        log.Info("Error when reading config")
    }
    //load from env variables
    replacer := strings.NewReplacer(".", "_")
    Viper.SetEnvKeyReplacer(replacer)
    Viper.AutomaticEnv()
}

func GetYamlValues() *yamlConfig {
    Db := &DBConfig{
        Server:        Viper.GetString("DATABASE.server"),
        Username:      Viper.GetString("DATABASE.username"),
        Password:      Viper.GetString("DATABASE.password"),
        Port:          Viper.GetString("DATABASE.port"),
        Schema:        Viper.GetString("DATABASE.schema"),
        MaxConnection: Viper.GetInt("DATABASE.connection_max"),
    }
    server := &ServerConfig{
        Port:  Viper.GetString("SERVICE.port"),
        Debug: Viper.GetString("SERVICE.debug"),
    }
    yml := &yamlConfig{*Db, *server}
    return yml
}