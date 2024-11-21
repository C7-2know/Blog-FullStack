package bootstrap

import (
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	APPEnv                 string `mapstructure:"APP_ENV"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBUri                  string `mapstructure:"MONGODB_URI"`
	AccessTokenExpiration  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiration int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure: "REFRESH_TOKEN_SECRET"`
	GeminiApikey         string `mapstructure:"GEMINI_API_KEY"`
	CloudName string `mapstructure:"CLOUD_NAME"`
	ApiKey  string `mapstructure:"CLOUD_API_KEY"`
	ApiSec  string `mapstructure:"CLOUD_API_SECRET"`


}

func NewEnv() *Env {
	env := Env{}
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	err := envMapToStruct(&env)
	if err != nil {
		log.Fatal(err)
	}

	if env.APPEnv=="development"{
		log.Println("Running in development mode")
	}
	return &env
}

func envMapToStruct(envStruct interface{}) error {
	structValue := reflect.ValueOf(envStruct).Elem()
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		envKey := field.Tag.Get("mapstructure")

		if envKey != "" {
			envValue := os.Getenv(envKey)
			fieldType := field.Type

			switch fieldType.Kind() {
			case reflect.String:
				structValue.Field(i).SetString(envValue)
			case reflect.Int:
				intValue, err := strconv.Atoi(envValue)
				if err != nil {
					return err
				}
				structValue.Field(i).SetInt(int64(intValue))
			}
		}
	}
	return nil

}
