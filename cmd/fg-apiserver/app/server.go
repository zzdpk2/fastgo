package app

import (
	"github.com/onexstack/fastgo/cmd/fg-apiserver/app/options"
	"github.com/onexstack/fastgo/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log/slog"
	"os"
)

var configFile string

// NewFastGOCommand to create a *cobra.Command project to initiate the program
func NewFastGOCommand() *cobra.Command {
	opts := options.NewServerOptions()
	cmd := &cobra.Command{
		Use:   "fg-apiserver",
		Short: "A very lightweight full go project",
		Long: `A very lightweight full go project, designed to help beginners quickly
		learn Go project development.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(),
		"Path to the fg-apiserver configuration file.")
	version.AddFlags(cmd.PersistentFlags())
	return cmd
}

// run function is the main running logic
func run(opts *options.ServerOptions) error {

	version.PrintAndExitIfRequested()

	// init slog
	initLog()

	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	if err := opts.Validate(); err != nil {
		return err
	}

	// Fetch the configurations of the program
	// Split the commandline configs and program configurations
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// Create an instance of the program
	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	// Start the server
	return server.Run()
}

func initLog() {
	// Log settings
	format := viper.GetString("log.format")
	level := viper.GetString("log.level")
	output := viper.GetString("log.output")

	var slevel slog.Level
	switch level {
	case "debug":
		slevel = slog.LevelDebug
	case "info":
		slevel = slog.LevelInfo
	case "warn":
		slevel = slog.LevelWarn
	case "error":
		slevel = slog.LevelError
	default:
		slevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: slevel}

	var w io.Writer
	var err error

	// Switch log output path
	switch output {
	case "":
		w = os.Stdout
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}

	// Reformat the log
	if err != nil {
		return
	}

	var handler slog.Handler
	switch format {
	case "json":
		handler = slog.NewJSONHandler(w, opts)
	case "text":
		handler = slog.NewTextHandler(w, opts)
	default:
		handler = slog.NewJSONHandler(w, opts)
	}

	// Set the log instance to a global scope
	slog.SetDefault(slog.New(handler))
}
