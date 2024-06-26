package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		Jwt      Jwt
		Kafka    Kafka
		Grpc     Grpc
		Paginate Paginate
	}

	App struct {
		Name  string
		Url   string
		stage string
	}

	Db struct {
		Url string
	}
	Jwt struct {
		AccessSecretKey  string
		RefreshSecretKey string
		ApiSecretKey     string
		AccessDuration   int64
		RefreshDuration  int64
		ApiDuration      int64
	}

	Kafka struct {
		Url    string
		Apikey string
		Secret string
	}

	Grpc struct {
		AuthUrl     string
		PlayerUrl   string
		ItemUrl     string
		IventoryUrl string
		PaymentUrl  string
	}

	Paginate struct {
		ItemNextPageBaseUrl      string
		InventoryNextPageBaseUrl string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
		Jwt: Jwt{
			AccessSecretKey:  os.Getenv("JWWT_ACCESS_SECRET_KEY"),
			RefreshSecretKey: os.Getenv("JWT_REFRESH_SECRET_KEY"),
			ApiSecretKey:     os.Getenv("JWT_API_SECRET_KEY"),
			AccessDuration:   loadENVInt64("JWT_ACCESS_DURATION"),
			RefreshDuration:  loadENVInt64("JWT_REFRESH_DURATION"),
		},
		Kafka: Kafka{
			Url:    os.Getenv("KAFKA_URL"),
			Apikey: os.Getenv("KAFKA_API_KEY"),
			Secret: os.Getenv("KAFKA_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl:     os.Getenv("GRPC_AUTH_URL"),
			PlayerUrl:   os.Getenv("GRPC_PLAYER_URL"),
			ItemUrl:     os.Getenv("GRPC_ITEM_URL"),
			IventoryUrl: os.Getenv("GRPC_INVENTORY_URL"),
			PaymentUrl:  os.Getenv("GRPC_PAYMENT_URL"),
		},
		Paginate: Paginate{
			ItemNextPageBaseUrl:      os.Getenv("PAGINATE_ITEM_NEXT_PAGE_BASED_URL"),
			InventoryNextPageBaseUrl: os.Getenv("PAGINATE_INVENTORY_NEXT_PAGE_BASED_URL"),
		},
	}
}

func loadENVInt64(envKey string) int64 {
	result, err := strconv.ParseInt(os.Getenv(envKey), 10, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
