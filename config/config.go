package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string // dev, test, prod
	Version     string

	GrpcPort string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	
	ProductServiceHost string
	ProductServicePort string

	DefaultOffset string
	DefaultLimit  string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "Order"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev"))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0.0"))

	config.GrpcPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":9001"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "your_db_password"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "your_db_name"))

	config.ProductServiceHost = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_HOST", "localhost"))
	config.ProductServicePort = cast.ToString(getOrReturnDefaultValue("PRODUCT_SERVICE_PORT", ":9001"))


	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
