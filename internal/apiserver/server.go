package apiserver

import (
	"context"
	"fmt"
	genericoptions "github.com/onexstack/fastgo/pkg/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

type Server struct {
	cfg *Config
}

func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func (s *Server) Run() error {
	fmt.Printf("Read MySQL host from config: %s\n", s.cfg.MySQLOptions.Addr)
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
