package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config estructura para almacenar la configuración del servicio IdP

type Config struct {
	Port       int    `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     int    `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
}

// LoadConfig carga la configuración desde variables de entorno o un archivo de configuración
func LoadConfig() Config {
	viper.AutomaticEnv()

	// Configuración por defecto
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "idp_db")
	viper.SetDefault("JWT_SECRET", "supersecretkey")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	return cfg
}
