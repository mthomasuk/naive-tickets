package handler

import (
	"fmt"
	"os"

	"github.com/mthomasuk/naive-tickets/internal/config"
	"github.com/mthomasuk/naive-tickets/internal/datastore"
	"github.com/stripe/stripe-go/client"
)

const (
	localCfgFile  = "./config/config.yaml"
	deployCfgFile = "/etc/signer-cover-sheet-creator/config/config.yaml"
)

var (
	cfg   config.Config
	sc    *client.API
	store *datastore.PgStore
)

func InitStripeAPI() {
	sc := &client.API{}
	sc.Init(cfg.Stripe.Key, nil)
}

func Init() {
	err := config.NewFromFile(deployCfgFile, &cfg)
	if err != nil {
		// Fall back to loading a local development config.yaml.
		err = config.NewFromFile(localCfgFile, &cfg)
		if err != nil {
			fmt.Printf("failed to load config :: %v", err)
			os.Exit(1)
		}
	}

	store, err = datastore.NewPostgres(cfg.Postgresql.Conn)
	if err != nil {
		fmt.Printf("failed to open connection to db :: %v", err)
		os.Exit(1)
	}

	InitStripeAPI()
}

func Run() {
	fmt.Println("ALL GOOD DUDE")
}
