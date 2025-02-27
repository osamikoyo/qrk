package app

import (
	"context"
	"github.com/osamikoyo/qrk/internal/config"
	"net/http"

	"github.com/osamikoyo/qrk/internal/transtort/handler"
	"github.com/osamikoyo/qrk/pkg/loger"
)

type App struct{
	loger loger.Logger
	server *http.Server
	mux *http.ServeMux
	assets string
}

func Init() App {
	cfg := config.New()

	return App{
		loger: loger.New(),
		server: &http.Server{
			Addr: cfg.Addr,
		},
		mux: http.NewServeMux(),
		assets: cfg.Assets,
	}
}

func (a App) Run(ctx context.Context) {
	go func ()  {
		<-ctx.Done()
		a.loger.Info().Msg("Server shutdowning!")
		a.server.Shutdown(ctx)
	}()

	handler.RegisterRoute(a.mux, handler.Deps{AssetsFs: http.Dir(a.assets)})

	a.loger.Info().Msg("Http server started!")

	if err := a.server.ListenAndServe();err != nil && err != http.ErrServerClosed{
		a.loger.Error().Err(err)
	}
}