package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mthomasuk/naive-tickets/internal/config"
	"github.com/mthomasuk/naive-tickets/internal/datastore"
	"github.com/mthomasuk/naive-tickets/internal/server"
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
	s     *http.Server
)

func InitServer() {
	s, err := srv.Init(cfg)
	if err != nil {
		fmt.Printf("failed to register server :: %v", err)
		os.Exit(1)
	}
	err = s.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to initialise server :: %v", err)
		os.Exit(1)
	}
}

func InitStripeAPI() {
	sc := &client.API{}
	sc.Init(cfg.Stripe.Key, nil)
}

func Init() {
	// Fall back to loading a local development config.yaml.
	err := config.NewFromFile(localCfgFile, &cfg)
	if err != nil {
		fmt.Printf("failed to load config :: %v", err)
		os.Exit(1)
	}

	store, err = datastore.NewPostgres(cfg.Postgresql.Conn)
	if err != nil {
		fmt.Printf("failed to open connection to db :: %v", err)
		os.Exit(1)
	}

	InitStripeAPI()
	InitServer()
}
