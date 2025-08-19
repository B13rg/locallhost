// Defines the cli-interface commands available to the user.
//
//nolint:gochecknoinits,gochecknoglobals

package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/b13rg/locallhost/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// exported version variable.
var version string

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "locallhost",
	Short: "Run locallhost server",
	Long:  `Start a server on a configured port that returns info about https requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		port := RootConfig.Port
		if port <= 0 {
			port = 8080
			log.Warn().Msgf("Invalid port num %d, defaulting to %d", RootConfig.Port, port)
		}

		logInterfaces(port)

		server.Serve(port)
	},
}

// Log interfaces tool will be serving on
func logInterfaces(port int) {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal().Err(err).Msg("error fetching interfaces")
	}
	log.Info().Msg("Serving locallhost on:")
	for _, intface := range interfaces {
		addrs, err := intface.Addrs()
		if err != nil {
			log.Warn().Err(err).Msg("error fetching interface addresses")
		}
		for _, addr := range addrs {
			if strings.Contains(addr.String(), ":") {
				// ipv6
				log.Info().Msgf("  http://[%v]:%d", strings.Split(addr.String(), "/")[0], port)
			} else {
				// ipv4
				log.Info().Msgf("  http://%v:%d", strings.Split(addr.String(), "/")[0], port)
			}
		}
	}
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ver string) {
	version = ver
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// Default options that are available to all commands.
type CmdRootOptions struct {
	// log more information about what the tool is doing. Overrides --loglevel
	Debug bool
	// set log level
	LogLevel string
	// enable colorized output (default true). Set to false to disable")
	Color bool
	// Profiling output directory.  Only captured if set.
	ProfilingDir string
	// CPU profiling output file handle.
	ProfilingCPUFile *os.File
	// HTTP port to listen on
	Port int
}

var RootConfig CmdRootOptions

func init() {
	// Ran before each command is ran
	cobra.OnInitialize(InitConfig, ProfilingInitializer)
	cobra.OnFinalize(ProfilingFinalizer)

	RootCmd.PersistentFlags().BoolVar(&RootConfig.Debug,
		"debug", false,
		"log additional information about what the tool is doing. Overrides --loglevel")
	RootCmd.PersistentFlags().StringVarP(&RootConfig.LogLevel,
		"loglevel", "L", "info",
		"set zerolog log level")
	RootCmd.PersistentFlags().BoolVar(&RootConfig.Color,
		"color", true,
		"enable colorized output")
	RootCmd.PersistentFlags().StringVarP(&RootConfig.ProfilingDir,
		"profiledir", "", "",
		"directory to write pprof profile data to")
	RootCmd.PersistentFlags().IntVarP(&RootConfig.Port,
		"port", "p", 8080,
		"Set http port for server to listen on",
	)
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	ConfigureLogger(RootConfig.Debug)

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.SetEnvPrefix("GCFG")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msg("Using config file:" + viper.ConfigFileUsed())
	}
}

// Stop profiling and write cpu and memory profiling files if configured.
func ProfilingFinalizer() {
	if RootConfig.ProfilingDir != "" {
		pprof.StopCPUProfile()
		if RootConfig.ProfilingCPUFile != nil {
			_ = RootConfig.ProfilingCPUFile.Close()
		}

		runtime.GC() // get up-to-date statistics

		// Various types of profiles that can be collected:
		// https://cs.opensource.google/go/go/+/go1.24.2:src/runtime/pprof/pprof.go;l=178
		var err error
		heapFile, err := os.Create(filepath.Join(RootConfig.ProfilingDir, "profile_heap.pb.gz"))
		if err != nil {
			log.Fatal().Err(err).Msg("could not write memory profile: ")
		}
		if err = pprof.WriteHeapProfile(heapFile); err != nil {
			_ = heapFile.Close()
			log.Fatal().Err(err).Msg("could not write memory profile: ")
		}
		_ = heapFile.Close()
	}
}

// Sets up program profiling.
func ProfilingInitializer() {
	var err error
	if RootConfig.ProfilingDir != "" {
		RootConfig.ProfilingCPUFile, err = os.Create(filepath.Join(RootConfig.ProfilingDir, "profile_cpu.pb.gz"))
		if err != nil {
			log.Fatal().Err(err).Msg("could not create CPU profile: ")
		}
		if err := pprof.StartCPUProfile(RootConfig.ProfilingCPUFile); err != nil {
			log.Fatal().Err(err).Msg("could not create CPU profile: ")
		}
	}
}
