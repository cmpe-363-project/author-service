package restapi

import (
	"author-service/internal/repository"
	"author-service/internal/restapi/routes"
	"author-service/pkg/logger"
	"net/http"
	"strconv"
)

type App struct {
	Version    string
	Logger     logger.Logger
	Repository repository.MysqlRepository

	Port int
	Host string
}

func (a *App) SetupAndRun() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/authors/{id}", routes.HandleGetAuthor(a.Logger, a.Repository))
	mux.HandleFunc("GET /api/version", routes.HandleGetVersion(a.Version))

	server := &http.Server{
		Addr:    a.Host + ":" + strconv.Itoa(a.Port),
		Handler: mux,
	}

	a.Logger.Info("Starting server", "host", a.Host, "port", strconv.Itoa(a.Port))
	if err := server.ListenAndServe(); err != nil {
		a.Logger.Error("Server failed to start")
		panic(err)
	}
}
