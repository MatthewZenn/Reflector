//go:generate goversioninfo -icon=Logo.ico

package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"github.com/spf13/cobra"
	"fmt"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

func init() {
	
	rootCmd.AddCommand(versionCmd)
  }
  
  var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
  }

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("%s", r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	log.Info().Msg("reflector-server is starting...")
	hub := newRouter()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to respond to WebSocket requests")
	}
}
