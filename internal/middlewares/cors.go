package middlewares

import (
	"github.com/iarsham/oauth2-example/configs"
	"github.com/rs/cors"
	"net/http"
)

func CorsMiddleware(cfg *configs.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: cfg.App.CorsOrigins,
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowCredentials: true,
		Debug:            cfg.App.Debug,
		MaxAge:           cfg.App.CorsMaxAge,
	})
}
