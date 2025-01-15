package main

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/laravel-ls/laravel-ls/laravel/providers/env"
	"github.com/laravel-ls/laravel-ls/laravel/providers/view"
	"github.com/laravel-ls/laravel-ls/lsp/server"
	"github.com/laravel-ls/laravel-ls/lsp/transport"
	"github.com/laravel-ls/laravel-ls/program"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var basePath string

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

var rootCmd = &cobra.Command{
	Use:     program.Name,
	Short:   "Language server for Laravel",
	Version: program.Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		basePath = expandHome(basePath)

		if err := os.MkdirAll(basePath, 0o755); err != nil {
			return err
		}

		logFilePath := path.Join(basePath, "logfile.log")

		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.SetLevel(log.DebugLevel)

		providerManager := provider.NewManager()
		providerManager.Add(view.NewProvider())
		providerManager.Add(env.NewProvider())

		defer treesitter.FreeQueryCache()

		log.Info("Starting laravel-ls")
		server := server.NewServer(providerManager)
		if err := server.Run(context.Background(), transport.NewStdio()); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&basePath, "basePath", "~/.local/laravel-ls", "base path")
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}` + "\n")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("Error")
		return
	}
}
