package config

import (
	"os"
	"time"

	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/spf13/viper"
)

// App config
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	MongoDB  MongoDB
	Cookie   Cookie
	Store    Store
	Session  Session
	Metrics  Metrics
	Logger   Logger
	AWS      AWS
	Jaeger   Jaeger
	Sentry   Sentry
}

// Server config
type ServerConfig struct {
	AppName           string
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbName   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleCons    int
	PoolSize       int
	PoolTimeout    int
	Passwordd      string
	DB             int
}

// MongoDB config
type MongoDB struct {
	MongoURI string
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// Store config
type Store struct {
	ImagesFolder string
}

// AWS S3
type AWS struct {
	Endpoint       string
	MinioAccessKey string
	MinioSecretKey string
	UseSSL         bool
	MinioEndpoint  string
}

// Jaeger
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

// Sentry
type Sentry struct {
	Dsn string
}

var globalConfig = Config{}

func GetConfig() Config {
	return globalConfig
}

func NewConfig(logger logger.Logger) *Config {
	configPath := utils.GetConfigPath(os.Getenv("config"))

	v := viper.New()
	v.SetConfigFile(configPath)
	// v.AddConfigPath("/config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Fatal("config file not found", err)
			return nil
		}
		logger.Fatal("cannot read configuration", err)
		return nil
	}

	err := v.Unmarshal(&globalConfig)
	if err != nil {
		logger.Fatal("environment cannot be loaded: ", err)
	}

	return &globalConfig
}
