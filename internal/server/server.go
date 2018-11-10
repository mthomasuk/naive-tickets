package srv

import (
	"net/http"
	"time"

	"github.com/mthomasuk/naive-tickets/internal/config"
	"github.com/mthomasuk/naive-tickets/internal/routes"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func Init(cfg config.Config) (*http.Server, error) {
	s := &http.Server{
		Addr:           cfg.Server.Port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/health", health)
	http.HandleFunc("/reserve", handler.Reserve)
	http.HandleFunc("/charge", handler.Charge)
	return s, nil
}
