package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Cors     CorsConfiguration
	Database DatabaseConfiguration
	Task     EnvironmentConfiguration
}
type MailConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}
type DeckersServerConfiguration struct {
	OAuthKey string `mapstructure:"oauthKey"`
	PostData string `mapstructure:"postData"`
}
type EnvironmentConfiguration struct {
	Server        ServerConfiguration        `mapstructure:"server"`
	Database      DatabaseConfiguration      `mapstructure:"database"`
	Mail          MailConfig                 `mapstructure:"email"`
	DeckersServer DeckersServerConfiguration `mapstructure:"deckersServer"`
	Tasks         TaskConfig                 `mapstructure:"tasks"`
	Cors          CorsConfiguration          `mapstructure:"cors"`
}
type TaskConfig struct {
	PostDeckerOutputsCron   string `mapstructure:"postDeckerOutputsCron"`
	PostDeckerOutputsFGCron string `mapstructure:"postDeckerOutputsFGCron"`
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

type CorsConfiguration struct {
	Global bool
	Ips    string
}

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	Sslmode  bool
	Logmode  bool
}

var Config *Configuration

func Setup(configPath string) error {
	var Configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	if err := viper.Unmarshal(&Configuration); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return err
	}

	Config = Configuration

	return nil
}

func GetConfig() *Configuration {
	return Config
}

// LoadEnv tải file .env và trả về giá trị biến môi trường
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}

// GetEnv lấy giá trị từ biến môi trường
func GetEnv(key string) string {
	return os.Getenv(key)
}
