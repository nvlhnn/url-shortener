package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database Database
	MemStore MemStore
	Limiter  Limiter
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type MemStore struct {
	Host     string
	Port     int
	Password string
	Database int
}

type Limiter struct {
	Max        int
	Expiration int
}

func NewConfig() *Config {
	c := NewConfigDefaults()

	// if err := c.loadConfigFromFile("config.toml"); err != nil {
	// 	log.Println("Unable to load config.toml, loaded defaults...")
	// }


	// map the env into Config struct
	c.Database.Host = os.Getenv("DATABASE_HOST")
	c.Database.Port = os.Getenv("DATABASE_PORT")
	c.Database.User = os.Getenv("DATABASE_USER")
	c.Database.Password = os.Getenv("DATABASE_PASSWORD")
	c.Database.Name = os.Getenv("DATABASE_NAME")

	c.MemStore.Host = os.Getenv("MEMSTORE_HOST")
	c.MemStore.Port, _ = strconv.Atoi(os.Getenv("MEMSTORE_PORT"))
	c.MemStore.Password = os.Getenv("MEMSTORE_PASSWORD")
	c.MemStore.Database, _ = strconv.Atoi(os.Getenv("MEMSTORE_DATABASE"))

	c.Limiter.Max, _ = strconv.Atoi(os.Getenv("LIMITER_MAX"))
	c.Limiter.Expiration, _ = strconv.Atoi(os.Getenv("LIMITER_EXPIRATION"))

	// c.applyEnvirontmentVariables()


	return c
}

func NewConfigDefaults() *Config {
	return &Config{}
}

// func (c *Config) loadConfigFromFile(path string) error {
// 	if _, err := toml.DecodeFile(path, c); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *Config) applyEnvirontmentVariables() {
// 	applyEnvirontmentVariable("DATABASE_HOST", &c.Database.Host)
// 	applyEnvirontmentVariable("DATABASE_PORT", &c.Database.Port)
// 	applyEnvirontmentVariable("DATABASE_USER", &c.Database.User)
// 	applyEnvirontmentVariable("DATABASE_PASSWORD", &c.Database.Password)
// 	applyEnvirontmentVariable("DATABASE_NAME", &c.Database.Name)

// 	applyEnvirontmentVariable("MEMSTORE_HOST", &c.MemStore.Host)
// 	applyEnvirontmentVariable("MEMSTORE_PORT", &c.MemStore.Port)
// 	applyEnvirontmentVariable("MEMSTORE_PASSWORD", &c.MemStore.Password)
// 	applyEnvirontmentVariable("MEMSTORE_DATABASE", &c.MemStore.Database)

// 	applyEnvirontmentVariable("LIMITER_MAX", &c.Limiter.Max)
// 	applyEnvirontmentVariable("LIMITER_EXPIRATION", &c.Limiter.Expiration)
// }

// func applyEnvirontmentVariable(key string, value interface{}) {
// 	if env, ok := os.LookupEnv(key); ok {
// 		switch v := value.(type) {
// 		case *string:
// 			*v = env
// 		case *bool:
// 			if env == "true" || env == "1" {
// 				*v = true
// 			} else if env == "false" || env == "0" {
// 				*v = false
// 			}
// 		case *int:
// 			if number, err := strconv.Atoi(env); err == nil {
// 				*v = number
// 			}
// 		}
// 	}
// }

// type Config struct {
// 	PORT          string `mapstructure:"PORT"`
// 	DB_URL        string `mapstructure:"DB_URL"`
// 	ENCODER_SCRET string `mapstructure:"ENCODER_SCRET"`
// }

// func LoadConfig() (config Config, err error) {
// 	viper.AddConfigPath("./pkg/config/envs")
// 	viper.SetConfigName("dev")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()

// 	if err != nil {
// 		return config, err
// 	}

// 	err = viper.Unmarshal(&config)

// 	return config, err
// }