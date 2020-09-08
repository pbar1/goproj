package cli

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pbar1/goproj/internal/pkg/api"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goproj",
	Short: "My default Go project starting point",
	Long:  `My default Go project starting point`,
	Run: func(cmd *cobra.Command, args []string) {
		s := api.NewServer(viper.GetInt("port"))
		err := s.Start()
		log.Fatal().Err(err)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	viper.Set("version", version)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().String("log-level", "info", "Log events >= this level. In order, one of: trace|debug|info|warn|error|fatal|panic")
	rootCmd.Flags().Bool("log-json", false, "Print logs in JSON format")
	rootCmd.Flags().Bool("log-line-numbers", true, "Log line numbers of the calling code")
	rootCmd.Flags().IntP("port", "p", 8080, "Local port for the HTTP API server to bind")

	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		log.Fatal().Err(err)
	}
}

func initConfig() {
	// configure logging
	logLevel, err := zerolog.ParseLevel(strings.ToLower(viper.GetString("log-level")))
	if err != nil {
		log.Warn().Err(err).Msg("unable to parse log level, using default: info")
	} else {
		zerolog.SetGlobalLevel(logLevel)
	}
	if !viper.GetBool("log-json") {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
	if viper.GetBool("log-line-numbers") {
		log.Logger = log.With().Caller().Logger()
	}
}
