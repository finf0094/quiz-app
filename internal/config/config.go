package config

import (
    "github.com/spf13/viper"
    "log"
)

type Config struct {
    Server struct {
        Port int `mapstructure:"port"`
    }
    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Name     string `mapstructure:"name"`
    }
}

var AppConfig Config

func LoadConfig() {
    // Загружаем значения по умолчанию
    viper.SetDefault("server.port", 8080)

    // Чтение конфигурации из файла
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")

    // Чтение переменных окружения
    viper.AutomaticEnv()
    viper.SetEnvPrefix("DB") // Префикс для переменных окружения

    // Привязка переменных окружения к полям конфигурации
    viper.BindEnv("database.host", "DB_HOST")
    viper.BindEnv("database.port", "DB_PORT")
    viper.BindEnv("database.user", "DB_USER")
    viper.BindEnv("database.password", "DB_PASSWORD")
    viper.BindEnv("database.name", "DB_NAME")

    // Чтение переменной окружения GIN_MODE для установки режима Gin
    viper.BindEnv("gin.mode", "GIN_MODE")

    // Чтение из файла, если он есть
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file: %s", err)
    }

    // Расшифровка значений в структуру конфигурации
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }
}
