package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/oswaldom-code/load-data-stock-donation-project/pkg/log"

	"github.com/spf13/viper"
)

const (
	// TODO: load PROJECT_NAME from environment variable
	PROJECT_NAME = "load-data-stock-donation-project"
	CONFIG_FILE  = "api"
)

type DBConfig struct {
	User               string
	Password           string
	Host               string
	Port               int
	Database           string
	MaxOpenConnections int
	SSLMode            string
	LogMode            string
	Engine             string
}

type ServerConfig struct {
	Host   string
	Port   string
	Scheme string
	Mode   string
	Static string
}

type LoggingConfig struct {
	Level        string
	ErrorLogFile string
}

type EnvironmentConfig struct {
	Environment string
}

type AuthConfig struct {
	Secret string
}

// ServerConfig validation
func (s *ServerConfig) Validate() error {
	if s.Host == "" || s.Port == "0" || s.Scheme == "" || s.Mode == "" {
		return fmt.Errorf("ServerConfig is invalid (host: %s, port: %s, scheme: %s, mode: %s)", s.Host, s.Port, s.Scheme, s.Mode)
	}
	return nil
}

// GetProjectPath returns the current project path
func GetProjectPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Warn("Warning, cannot get current path")
		return ""
	}
	// Traverse back from current directory until service base dir is reach and add to config path
	for !strings.HasSuffix(dir, PROJECT_NAME) && dir != "/" {
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			break
		}
	}
	return dir
}

func Load() {
	viper.SetConfigName(CONFIG_FILE)
	viper.SetEnvPrefix(PROJECT_NAME)
	viper.AddConfigPath(GetProjectPath() + "/config")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_", ".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config error:", err.Error())
	}
	log.SetLogLevel(GetLogConfig().Level)
}

func GetDBConfig() DBConfig {
	env := GetEnvironmentConfig().Environment
	config := DBConfig{
		User:               viper.GetString(env + ".db.user"),
		Password:           viper.GetString(env + ".db.password"),
		Host:               viper.GetString(env + ".db.host"),
		Port:               viper.GetInt(env + ".db.port"),
		Database:           viper.GetString(env + ".db.database"),
		MaxOpenConnections: viper.GetInt(env + ".db.max_connections"),
		SSLMode:            viper.GetString(env + ".db.ssl_mode"),
		LogMode:            viper.GetString(env + ".db.log_mode"),
		Engine:             viper.GetString(env + ".db.engine"),
	}
	log.DebugWithFields("DBConfig", log.Fields{"config": config})
	return config
}

func GetServerConfig() ServerConfig {
	env := GetEnvironmentConfig().Environment
	config := ServerConfig{
		Host:   viper.GetString(env + ".server.host"),
		Port:   viper.GetString(env + ".server.port"),
		Scheme: viper.GetString(env + ".server.scheme"),
		Mode:   viper.GetString(env + ".server.mode"),
		Static: viper.GetString(env + ".server.static"),
	}
	log.DebugWithFields("ServerConfig", log.Fields{"config": config})
	return config
}

func GetStaticPath() string {
	env := GetEnvironmentConfig().Environment
	return viper.GetString(env + ".server.static")
}

func GetLogConfig() LoggingConfig {
	env := GetEnvironmentConfig().Environment
	return LoggingConfig{
		Level:        viper.GetString(env + ".log.level"),
		ErrorLogFile: viper.GetString(env + ".log.errorLogFile"),
	}
}

func GetEnvironmentConfig() EnvironmentConfig {
	return EnvironmentConfig{
		Environment: viper.GetString("environment"),
	}
}

func GetAuthConfig() AuthConfig {
	return AuthConfig{
		Secret: viper.GetString(GetEnvironmentConfig().Environment + ".auth.secret"),
	}
}

func (s ServerConfig) AsUri() string {
	return s.Host + ":" + s.Port
}

func (s DBConfig) GetConnectionString() string {
	bStr, err := json.Marshal(s)
	if err != nil {
		log.FatalWithFields("Cannot marshal DSN", log.Fields{"err": err, "dsn": s})
		return ""
	}
	connectionString := string(bStr)
	connectionString = strings.NewReplacer(":", "=", "{", "", "}", "", `"`, "", ",", " ").Replace(connectionString)
	return connectionString
}
