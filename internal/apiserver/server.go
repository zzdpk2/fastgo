package apiserver

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/onexstack/fastgo/internal/pkg/core"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	mw "github.com/onexstack/fastgo/internal/pkg/middleware"
	genericoptions "github.com/onexstack/fastgo/pkg/options"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	engine := gin.New()

	// the middleware gin.Recovery() to catch panic and have it recover
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}
	engine.Use(mws...)
	// Register 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errorsx.ErrNotFound.WithMessage("Page not found"), nil)
	})

	// Register /healthz handler
	engine.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	// Create HTTP Server instance
	httpsrv := &http.Server{
		Addr: cfg.Addr, Handler: engine,
	}

	return &Server{cfg: cfg, srv: httpsrv}, nil

}

func (s *Server) Run() error {
	// fmt.Printf("Read MySQL host from config: %s\n", s.cfg.MySQLOptions.Addr)
	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	// Block to prevent the program from quitting
	// select {}
	// block until receive Ctrl+C / kill（TERM）
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // Won't be judged as a deadlock
	// Elegantly shutdown service
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the dependant service, then independent service
	// shut the service in 10s, then quit
	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Insecure Server forced to shutdown", "err", err)
		return err
	}

	slog.Info("Server exited")
	return nil
}
