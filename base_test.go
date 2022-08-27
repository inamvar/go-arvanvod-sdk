package arvanvod_test

import (
	"os"

	"github.com/inamvar/go-arvanvod-sdk"
	"github.com/joho/godotenv"
)

func getClient() *arvanvod.Client {
	godotenv.Load()
	opt := &arvanvod.ClientOptions{
		ApiKey:  os.Getenv("arvan_api_key"),
		BaseUrl: os.Getenv("arvan_base_url"),
	}
	client := arvanvod.NewClient(opt)

	return client
}
