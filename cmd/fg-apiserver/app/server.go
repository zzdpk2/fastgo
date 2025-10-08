package app

import (
	"github.com/onexstack/fastgo/cmd/fg-apiserver/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	return cmd
}

// run function is the main running logic
func run(opts *options.ServerOptions) error {
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
