package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/laravel-ls/laravel-ls/config"
	appProvider "github.com/laravel-ls/laravel-ls/laravel/providers/app"
	assetsProvider "github.com/laravel-ls/laravel-ls/laravel/providers/assets"
	configProvider "github.com/laravel-ls/laravel-ls/laravel/providers/config"
	envProvider "github.com/laravel-ls/laravel-ls/laravel/providers/env"
	viewProvider "github.com/laravel-ls/laravel-ls/laravel/providers/view"
	"github.com/laravel-ls/laravel-ls/lsp/server"
	"github.com/laravel-ls/laravel-ls/lsp/transport"
	"github.com/laravel-ls/laravel-ls/program"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	basePath string
	logFile  *os.File
)

// Expand "~" to user's home directory.
func expandHome(path string) string {
	if p, found := strings.CutPrefix(path, "~"); found {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return homedir + p
	}
	return path
}

func openLogFile(filename string) (*os.File, error) {
	logFilePath := path.Join(basePath, filename)
	return os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
}

func run(cmd *cobra.Command, args []string) error {
	basePath = expandHome(basePath)

	// Setup config
	viper.AddConfigPath(".")
	viper.AddConfigPath(basePath)
	err := viper.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}

	cfg, err := config.Parse(viper.GetViper())
	if err != nil {
		return err
	}

	if err := os.MkdirAll(basePath, 0o755); err != nil {
		return err
	}

	logFile, err = openLogFile(cfg.Log.Filename)
	if err != nil {
		return err
	}

	log.SetOutput(logFile)
	log.SetLevel(cfg.Log.Level)

	if len(viper.ConfigFileUsed()) > 0 {
		log.WithField("file", viper.ConfigFileUsed()).Debug("config file used.")
	} else {
		log.Debug("no config file used.")
	}

	providerManager := provider.NewManager(
		viewProvider.NewProvider(),
		envProvider.NewProvider(),
		assetsProvider.NewProvider(),
		appProvider.NewProvider(),
		configProvider.NewProvider(),
	)

	defer treesitter.FreeQueryCache()

	log.Info("Starting laravel-ls")
	server := server.NewServer(providerManager)
	if err := server.Run(context.Background(), transport.NewStdio()); err != nil {
		return err
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func bindFlagsToConfig(flags *pflag.FlagSet) {
	must(viper.BindPFlag("log.filename", flags.Lookup("log")))
	must(viper.BindPFlag("log.level", flags.Lookup("log-level")))
}

func Run() {
	cmd := cobra.Command{
		Use:     program.Name,
		Short:   "Language server for Laravel",
		Version: program.Version(),
		RunE:    run,
	}

	cmd.PersistentFlags().StringVar(&basePath, "basePath", "~/.local/laravel-ls", "base path")
	cmd.PersistentFlags().String("log", "log", "Log file, relative to basePath")
	cmd.PersistentFlags().String("log-level", "info", fmt.Sprintf("Logging level, one of: %v", log.AllLevels))
	cmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}` + "\n")

	bindFlagsToConfig(cmd.PersistentFlags())

	if err := cmd.Execute(); err != nil {
		log.WithError(err).Error("Error")
	}

	if logFile != nil {
		logFile.Close()
	}
}
