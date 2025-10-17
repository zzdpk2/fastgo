package apiserver

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	genericoptions "github.com/onexstack/fastgo/pkg/options"
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

	// Register 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "PageNotFound", "message": "Page not found."})
	})

	// Register /healthz handler
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
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
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	// Block to prevent the program from quitting
	// select {}
	// block until receive Ctrl+C / kill（TERM）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done() // Won't be judged as a deadlock
	// To do resources collection like shutting down DB etc, sleep is just a mock
	time.Sleep(100 * time.Millisecond)
	return nil
}
